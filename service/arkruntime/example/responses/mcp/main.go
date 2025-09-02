package main

import (
	"context"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"io"
	"os"
)

/**
 * Authentication
 * If you authorize your endpoint using an API key, you can set your api key to environment variable "ARK_API_KEY"
 * client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
 * Note: If you use an API key, this API key will not be refreshed.
 * To prevent the API from expiring and failing after some time, choose an API key with no expiration date.
 */

func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()

	fmt.Println("----- round 1 message -----")
	// round 1 message
	serverDescription := "test desc"
	inputMessage := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: "查一下 mark3labs/mcp-go 这个仓库的结构",
					},
				},
			},
		},
	}
	createResponsesReq := &responses.ResponsesRequest{
		Model: "doubao-seed-1-6",
		Input: &responses.ResponsesInput{
			Union: &responses.ResponsesInput_ListValue{
				ListValue: &responses.InputItemList{ListValue: []*responses.InputItem{{
					Union: &responses.InputItem_InputMessage{
						InputMessage: inputMessage,
					},
				}}},
			},
		},
		//Caching: &responses.ResponsesCaching{Type: responses.CacheType_enabled.Enum()},
		Tools: []*responses.ResponsesTool{
			{
				Union: &responses.ResponsesTool_ToolMcp{
					ToolMcp: &responses.ToolMcp{
						Type:              responses.ToolType_mcp,
						ServerLabel:       "deepwiki",
						ServerDescription: &serverDescription,
						ServerUrl:         "https://mcp.deepwiki.com/mcp",
					},
				},
			},
		},
	}

	resp, err := client.CreateResponsesStream(ctx, createResponsesReq, arkruntime.WithCustomHeader("ark-beta-mcp", "true"))
	if err != nil {
		fmt.Printf("stream error: %v\n", err)
		return
	}
	var responseId string
	var approvalRequestId string
	for {
		event, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("stream error: %v\n", err)
			return
		}
		handleEvent(event)
		if responseEvent := event.GetResponse(); responseEvent != nil {
			responseId = responseEvent.GetResponse().GetId()
		}
		if event.GetItem().GetItem().GetFunctionMcpApprovalRequest() != nil {
			approvalRequestId = event.GetItem().GetItem().GetFunctionMcpApprovalRequest().GetId()
		}
	}

	fmt.Println()
	fmt.Println("-----round 2---------")
	createResponsesReq2 := &responses.ResponsesRequest{
		Model: "doubao-seed-1-6",
		Input: &responses.ResponsesInput{
			Union: &responses.ResponsesInput_ListValue{
				ListValue: &responses.InputItemList{ListValue: []*responses.InputItem{{
					Union: &responses.InputItem_McpApprovalResponse{
						McpApprovalResponse: &responses.ItemFunctionMcpApprovalResponse{
							Type:              responses.ItemType_mcp_approval_response,
							ApprovalRequestId: approvalRequestId,
							Approve:           true,
						},
					},
				}}},
			},
		},
		//Caching: &responses.ResponsesCaching{Type: responses.CacheType_enabled.Enum()},
		Tools: []*responses.ResponsesTool{
			{
				Union: &responses.ResponsesTool_ToolMcp{
					ToolMcp: &responses.ToolMcp{
						Type:              responses.ToolType_mcp,
						ServerLabel:       "deepwiki",
						ServerDescription: &serverDescription,
						ServerUrl:         "https://mcp.deepwiki.com/mcp",
					},
				},
			},
		},
		PreviousResponseId: &responseId,
	}
	resp2, err := client.CreateResponsesStream(ctx, createResponsesReq2, arkruntime.WithCustomHeader("ark-beta-mcp", "true"))
	if err != nil {
		fmt.Printf("stream error: %v\n", err)
		return
	}
	for {
		event, err := resp2.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("stream error: %v\n", err)
			return
		}
		handleEvent(event)
	}

	fmt.Println("----- round 3 never require approval -----")
	createResponsesReq3 := &responses.ResponsesRequest{
		Model: "doubao-seed-1-6",
		Input: &responses.ResponsesInput{
			Union: &responses.ResponsesInput_ListValue{
				ListValue: &responses.InputItemList{ListValue: []*responses.InputItem{{
					Union: &responses.InputItem_InputMessage{
						InputMessage: inputMessage,
					},
				}}},
			},
		},
		Tools: []*responses.ResponsesTool{
			{
				Union: &responses.ResponsesTool_ToolMcp{
					ToolMcp: &responses.ToolMcp{
						Type:              responses.ToolType_mcp,
						ServerLabel:       "deepwiki",
						ServerDescription: &serverDescription,
						ServerUrl:         "https://mcp.deepwiki.com/mcp",
						RequireApproval: &responses.McpRequireApproval{
							Union: &responses.McpRequireApproval_Mode{
								Mode: responses.ApprovalMode_never,
							},
						},
					},
				},
			},
		},
	}

	resp3, err := client.CreateResponsesStream(ctx, createResponsesReq3, arkruntime.WithCustomHeader("ark-beta-mcp", "true"))
	if err != nil {
		fmt.Printf("stream error: %v\n", err)
		return
	}
	for {
		event, err := resp3.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("stream error: %v\n", err)
			return
		}
		handleEvent(event)
	}
}

func handleEvent(event *responses.Event) {
	switch event.GetEventType() {
	case responses.EventType_response_reasoning_summary_text_delta.String():
		print(event.GetReasoningText().GetDelta())
	case responses.EventType_response_reasoning_summary_text_done.String(): // aggregated reasoning text
		fmt.Printf("\naggregated reasoning text: %s\n", event.GetReasoningText().GetText())
	case responses.EventType_response_output_text_delta.String():
		print(event.GetText().GetDelta())
	case responses.EventType_response_output_text_done.String(): // aggregated output text
		fmt.Printf("\naggregated output text: %s\n", event.GetText().GetText())
	case responses.EventType_response_output_item_done.String():
		if event.GetItem().GetItem().GetFunctionMcpApprovalRequest() != nil {
			fmt.Printf("\nmcp call: %s; arguments %s\n", event.GetItem().GetItem().GetFunctionMcpApprovalRequest().GetName(), event.GetItem().GetItem().GetFunctionMcpApprovalRequest().GetArguments())
		}
	case responses.EventType_response_mcp_call_arguments_delta.String():
		fmt.Printf("\nmcp call arguments: %s\n", event.GetResponseMcpCallArgumentsDelta().GetDelta())
	case responses.EventType_response_mcp_call_completed.String():
		fmt.Printf("\nmcp call completed.\n")
	case responses.EventType_response_mcp_list_tools_in_progress.String():
		fmt.Printf("\nlisting mcp tools.\n")
	case responses.EventType_response_mcp_list_tools_completed.String():
		fmt.Printf("\nDone listing mcp tools.\n")
	default:
		return
	}
}

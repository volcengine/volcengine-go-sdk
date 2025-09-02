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
	maxToolCalls := int64(1)

	inputMessage := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: "查一下今天的AI相关新闻",
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
		Tools: []*responses.ResponsesTool{
			{
				Union: &responses.ResponsesTool_ToolWebSearch{
					ToolWebSearch: &responses.ToolWebSearch{
						Type: responses.ToolType_web_search,
					},
				},
			},
		},
		MaxToolCalls: &maxToolCalls,
	}

	resp, err := client.CreateResponsesStream(ctx, createResponsesReq)
	if err != nil {
		fmt.Printf("stream error: %v\n", err)
		return
	}

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
	}

	fmt.Println()
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

package basic
package main

import (
	"context"
	"fmt"
	"io"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

/**
 * Authentication
 * If you authorize your endpoint using an API key, you can set your api key to environment variable "ARK_API_KEY"
 * client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
 * Note: If you use an API key, this API key will not be refreshed.
 * To prevent the API from expiring and failing after some time, choose an API key with no expiration date.
 */

func main() {
	client := arkruntime.NewClientWithApiKey("01c254dd-76c9-4b0a-a717-cea9330cdfad", arkruntime.WithBaseUrl("https://ark-stg.cn-beijing.volces.com/api/v3/"))
	ctx := context.Background()

	fmt.Println("----- round 1 message -----")
	// round 1 message
	serverDescription := "test desc"
	inputMessage := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Image{
					Image: &responses.ContentItemImage{
						Type:     responses.ContentItemType_input_image,
						Detail:   responses.ContentItemImageDetail_auto.Enum(),
						ImageUrl: "https://ark-project.tos-cn-beijing.volces.com/images/view.jpeg",
					},
				},
			},
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: "请给出图片的描述",
					},
				},
			},
		},
	}
	createResponsesReq := &responses.ResponsesRequest{
		Model: "ep-20250826102747-62tcl",
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

	resp, err := client.CreateResponsesStream(ctx, createResponsesReq)
	if err != nil {
		fmt.Printf("stream error: %v\n", err)
		return
	}
	var responseId string
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
	}

	fmt.Println()
	// round 2 messages
	fmt.Println("-----round 2---------")
	createResponsesReq.Input.Union = &responses.ResponsesInput_StringValue{StringValue: "上述对话中有几幅图片，每幅图片的描述是？"}
	createResponsesReq.PreviousResponseId = volcengine.String(responseId)
	resp, err = client.CreateResponsesStream(ctx, createResponsesReq)
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
		if responseEvent := event.GetResponse(); responseEvent != nil {
			responseId = responseEvent.GetResponse().GetId()
		}
	}
	fmt.Println(responseId)
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
	default:
		print(event.GetEventType())
		return
	}
}

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/samber/lo"
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
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()

	fmt.Println("----- round 1 message -----")
	// round 1 message
	inputMessage := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Image{
					Image: &responses.ContentItemImage{
						Type:     responses.ContentItemType_input_image,
						Detail:   responses.ContentItemImageDetail_auto.Enum(),
						ImageUrl: lo.ToPtr("https://ark-project.tos-cn-beijing.volces.com/images/view.jpeg"),
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
		Caching: &responses.ResponsesCaching{Type: responses.CacheType_enabled.Enum()},
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
		return
	}
}

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
)

/**
 * Authentication
 * If you authorize your endpoint using an API key, you can set your api key to environment variable "ARK_API_KEY"
 * client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
 * Note: If you use an API key, this API key will not be refreshed.
 * To prevent the API from expiring and failing after some time, choose an API key with no expiration date.
 */

const (
	ChatFeature            = "chat"
	DeepChatFeature        = "deepChat"
	AiSearchFeature        = "aiSearch"
	ReasoningSearchFeature = "reasoningSearch"
)

var enabeld = responses.ResponseDoubaoAppFeatureType_enabled

// var optionalRoleDescription = "你是小明"
var setting = responses.DoubaoAppFeatureItem{
	Type: &enabeld,
	//RoleDescription: &optionalRoleDescription,
}

var featureToQuery = map[string]string{
	ChatFeature:            "你好介绍一下你自己",
	DeepChatFeature:        "为什么天空是蓝色",
	AiSearchFeature:        "今天的AI新闻",
	ReasoningSearchFeature: "今天的AI新闻",
}

var featureToToolParam = map[string]*responses.DoubaoAppFeature{
	ChatFeature: {
		Chat: &setting,
	},
	DeepChatFeature: {
		DeepChat: &setting,
	},
	AiSearchFeature: {
		AiSearch: &setting,
	},
	ReasoningSearchFeature: {
		ReasoningSearch: &setting,
	},
}

func main() {
	stream(ChatFeature)    // change to DeepChatFeature, AiSearchFeature, ReasoningSearchFeature to test different features
	nonStream(ChatFeature) // change to DeepChatFeature, AiSearchFeature, ReasoningSearchFeature to test different features
}

func nonStream(feature string) {
	fmt.Println("non stream")
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()

	fmt.Println("----- round 1 message -----")
	// round 1 message
	inputMessage := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: featureToQuery[feature],
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
				Union: &responses.ResponsesTool_ToolDoubaoApp{
					ToolDoubaoApp: &responses.ToolDoubaoApp{
						Type:    responses.ToolType_doubao_app,
						Feature: featureToToolParam[feature],
					},
				},
			},
		},
	}

	resp, err := client.CreateResponses(ctx, createResponsesReq, arkruntime.WithCustomHeader("ark-beta-doubao-app", "true"))
	if err != nil {
		fmt.Printf("request error: %v\n", err)
		return
	}
	var responseId string
	fmt.Printf("response: %v\n", resp)
	responseId = resp.Id

	fmt.Println()
	fmt.Println("-----round 2---------")
	inputMessage2 := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: "刚刚我们聊了什么",
					},
				},
			},
		},
	}
	createResponsesReq2 := &responses.ResponsesRequest{
		Model: "doubao-seed-1-6",
		Input: &responses.ResponsesInput{
			Union: &responses.ResponsesInput_ListValue{
				ListValue: &responses.InputItemList{ListValue: []*responses.InputItem{{
					Union: &responses.InputItem_InputMessage{
						InputMessage: inputMessage2,
					},
				}}},
			},
		},
		Tools: []*responses.ResponsesTool{
			{
				Union: &responses.ResponsesTool_ToolDoubaoApp{
					ToolDoubaoApp: &responses.ToolDoubaoApp{
						Type:    responses.ToolType_doubao_app,
						Feature: featureToToolParam[feature],
					},
				},
			},
		},
		PreviousResponseId: &responseId,
	}
	resp2, err := client.CreateResponses(ctx, createResponsesReq2, arkruntime.WithCustomHeader("ark-beta-doubao-app", "true"))
	if err != nil {
		fmt.Printf("request error: %v\n", err)
		return
	}
	fmt.Printf("response: %v\n", resp2)
}

func stream(feature string) {
	fmt.Println("stream")

	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()
	fmt.Println("----- round 1 message -----")
	// round 1 message
	inputMessage := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: featureToQuery[feature],
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
				Union: &responses.ResponsesTool_ToolDoubaoApp{
					ToolDoubaoApp: &responses.ToolDoubaoApp{
						Type:    responses.ToolType_doubao_app,
						Feature: featureToToolParam[feature],
					},
				},
			},
		},
	}

	resp, err := client.CreateResponsesStream(ctx, createResponsesReq, arkruntime.WithCustomHeader("ark-beta-doubao-app", "true"))
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
	fmt.Println("-----round 2---------")
	inputMessage2 := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: "刚刚我们聊了什么",
					},
				},
			},
		},
	}
	createResponsesReq2 := &responses.ResponsesRequest{
		Model: "doubao-seed-1-6",
		Input: &responses.ResponsesInput{
			Union: &responses.ResponsesInput_ListValue{
				ListValue: &responses.InputItemList{ListValue: []*responses.InputItem{{
					Union: &responses.InputItem_InputMessage{
						InputMessage: inputMessage2,
					},
				}}},
			},
		},
		Tools: []*responses.ResponsesTool{
			{
				Union: &responses.ResponsesTool_ToolDoubaoApp{
					ToolDoubaoApp: &responses.ToolDoubaoApp{
						Type:    responses.ToolType_doubao_app,
						Feature: featureToToolParam[feature],
					},
				},
			},
		},
		PreviousResponseId: &responseId,
	}
	resp2, err := client.CreateResponsesStream(ctx, createResponsesReq2, arkruntime.WithCustomHeader("ark-beta-doubao-app", "true"))
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
}

func handleEvent(event *responses.Event) {
	switch event.GetEventType() {
	case responses.EventType_response_doubao_app_call_search_searching.String():
		print(event.GetResponseDoubaoAppCallSearchSearching().GetSearchingState())
	case responses.EventType_response_doubao_app_call_search_completed.String():
		print(event.GetResponseDoubaoAppCallSearchCompleted().GetSummary())
	case responses.EventType_response_doubao_app_call_reasoning_search_completed.String():
		print(event.GetResponseDoubaoAppCallReasoningSearchCompleted().GetSummary())
	case responses.EventType_response_doubao_app_call_output_text_delta.String():
		print(event.GetResponseDoubaoAppCallOutputTextDelta().GetDelta())
	case responses.EventType_response_doubao_app_call_output_text_done.String():
		fmt.Printf("\naggregated output text: %s\n", event.GetResponseDoubaoAppCallOutputTextDone().GetText())
	case responses.EventType_response_doubao_app_call_reasoning_text_delta.String():
		print(event.GetResponseDoubaoAppCallReasoningTextDelta().GetDelta())
	case responses.EventType_response_doubao_app_call_reasoning_text_done.String():
		fmt.Printf("\naggregated reasoning text: %s\n", event.GetResponseDoubaoAppCallReasoningTextDone().GetText())
	default:
		return
	}
}

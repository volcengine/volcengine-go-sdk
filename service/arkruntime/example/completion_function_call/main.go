package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

/**
 * Authentication
 * 1.If you authorize your endpoint using an API key, you can set your api key to environment variable "ARK_API_KEY"
 * client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
 * Note: If you use an API key, this API key will not be refreshed.
 * To prevent the API from expiring and failing after some time, choose an API key with no expiration date.
 *
 * 2.If you authorize your endpoint with Volcengine Identity and Access Management（IAM), set your api key to environment variable "VOLC_ACCESSKEY", "VOLC_SECRETKEY"
 * client := arkruntime.NewClientWithAkSk(os.Getenv("VOLC_ACCESSKEY"), os.Getenv("VOLC_SECRETKEY"))
 * To get your ak&sk, please refer to this document(https://www.volcengine.com/docs/6291/65568)
 * For more information，please check this document（https://www.volcengine.com/docs/82379/1263279）
 */

func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()

	fmt.Println("----- function call request -----")
	req := model.CreateChatCompletionRequest{
		Model: "${YOUR_ENDPOINT_ID}",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String("What's the weather like in Boston today?"),
				},
			},
		},
		Tools: []*model.Tool{
			{
				Type: model.ToolTypeFunction,
				Function: &model.FunctionDefinition{
					Name:        "get_current_weather",
					Description: "Get the current weather in a given location",
					Parameters: map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"location": map[string]interface{}{
								"type":        "string",
								"description": "The city and state, e.g. San Francisco, CA",
							},
							"unit": map[string]interface{}{
								"type": "string",
								"enum": []string{"celsius", "fahrenheit"},
							},
						},
						"required": []string{"location"},
					},
				},
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("standard chat error: %v\n", err)
		return
	}

	s, _ := json.Marshal(resp)
	fmt.Println(string(s))

	fmt.Println("----- function call stream request -----")
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("stream chat error: %v\n", err)
		return
	}
	defer stream.Close()

	finalToolCalls := make(map[int]*model.ToolCall)
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Stream chat error: %v\n", err)
			break
		}
		fmt.Print(recv.Choices[0].Delta.Content)

		for _, toolCall := range recv.Choices[0].Delta.ToolCalls {
			index := 0
			if toolCall.Index != nil {
				index = *toolCall.Index
			}

			if _, exists := finalToolCalls[index]; !exists {
				finalToolCalls[index] = toolCall
			}

			finalToolCalls[index].Function.Arguments += toolCall.Function.Arguments
		}
	}
	for _, finalToolCall := range finalToolCalls {
		fmt.Println(*finalToolCall)
	}
}

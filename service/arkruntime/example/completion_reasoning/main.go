package main

import (
	"context"
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

	fmt.Println("----- streaming request -----")
	req := model.CreateChatCompletionRequest{
		Model: "${YOUR_ENDPOINT_ID}",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String("How many Rs are there in the word 'strawberry'?"),
				},
			},
		},
		Thinking: &model.Thinking{
			Type: model.ThinkingTypeEnabled,
		},
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("stream chat error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Stream chat error: %v\n", err)
			break
		}

		if len(recv.Choices) > 0 {
			if recv.Choices[0].Delta.ReasoningContent != nil && *recv.Choices[0].Delta.ReasoningContent != "" {
				fmt.Print(*recv.Choices[0].Delta.ReasoningContent)
			} else {
				fmt.Print(recv.Choices[0].Delta.Content)
			}
		}
	}
	fmt.Println()

	fmt.Println("----- standard request -----")
	req = model.CreateChatCompletionRequest{
		Model: "${YOUR_ENDPOINT_ID}",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String("How many Rs are there in the word 'strawberry'?"),
				},
			},
		},
		Thinking: &model.Thinking{
			Type: model.ThinkingTypeEnabled,
		},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("standard chat error: %v\n", err)
		return
	}
	if resp.Choices[0].Message.ReasoningContent != nil {
		fmt.Println(*resp.Choices[0].Message.ReasoningContent)
	}
	fmt.Println(*resp.Choices[0].Message.Content.StringValue)
}

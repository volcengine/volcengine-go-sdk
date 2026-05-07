package main

import (
	"context"
	"fmt"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()

	fmt.Println("----- audio input (url) -----")
	req := model.CreateChatCompletionRequest{
		Model: "${YOUR_ENDPOINT_ID}",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					ListValue: []*model.ChatCompletionMessageContentPart{
						{
							Type: model.ChatCompletionMessageContentPartTypeText,
							Text: "你好，这个音频里说了什么",
						},
						{
							Type: model.ChatCompletionMessageContentPartTypeAudioURL,
							InputAudio: &model.ChatMessageAudioURL{
								URL: "https://222lytest.tos-cn-beijing.volces.com/question.wav",
							},
						},
					},
				},
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("audio chat error: %v\n", err)
		return
	}

	if len(resp.Choices) > 0 && resp.Choices[0].Message.Content != nil &&
		resp.Choices[0].Message.Content.StringValue != nil {
		fmt.Println("response:", *resp.Choices[0].Message.Content.StringValue)
	}

	u := resp.Usage
	fmt.Println("----- usage -----")
	fmt.Printf("prompt_tokens:     %d\n", u.PromptTokens)
	fmt.Printf("completion_tokens: %d\n", u.CompletionTokens)
	fmt.Printf("total_tokens:      %d\n", u.TotalTokens)
	if u.PromptTokensDetails.AudioTokens != nil {
		fmt.Printf("prompt audio_tokens:     %d\n", *u.PromptTokensDetails.AudioTokens)
	} else {
		fmt.Println("prompt audio_tokens:     <nil>")
	}
	if u.CompletionTokensDetails.AudioTokens != nil {
		fmt.Printf("completion audio_tokens: %d\n", *u.CompletionTokensDetails.AudioTokens)
	} else {
		fmt.Println("completion audio_tokens: <nil>")
	}

	if u.PromptTokens == 0 || u.CompletionTokens == 0 {
		fmt.Println("WARN: token usage looks empty — audio input may not have been processed")
		os.Exit(1)
	}
	if u.PromptTokensDetails.AudioTokens == nil || *u.PromptTokensDetails.AudioTokens == 0 {
		fmt.Println("WARN: no prompt audio_tokens reported — input_audio likely not recognized")
		os.Exit(1)
	}
	fmt.Println("OK: audio input accepted and accounted for in usage")
}

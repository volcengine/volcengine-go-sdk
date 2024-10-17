package main

import (
	"context"
	"fmt"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
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

	fmt.Println("----- image input -----")
	req := model.ChatCompletionRequest{
		Model: "${YOUR_ENDPOINT_ID}",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					ListValue: []*model.ChatCompletionMessageContentPart{
						{
							Type: model.ChatCompletionMessageContentPartTypeText,
							Text: "这是哪里？",
						},
						{
							Type: model.ChatCompletionMessageContentPartTypeImageURL,
							ImageURL: &model.ChatMessageImageURL{
								URL: "https://ark-project.tos-cn-beijing.ivolces.com/images/view.jpeg",
							},
						},
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
	fmt.Println(*resp.Choices[0].Message.Content.StringValue)
}

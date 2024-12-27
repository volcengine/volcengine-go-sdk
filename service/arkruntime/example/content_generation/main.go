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
 */
func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()

	fmt.Println("----- create content generation task -----")
	createReq := model.CreateContentGenerationTaskRequest{
		Model: "{YOUR_ENDPOINT_ID}", // Replace with your endpoint ID
		Content: []*model.CreateContentGenerationContentItem{
			{
				Type: model.ContentGenerationContentItemTypeText,
				Text: "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原 --ratio 1:1",
			},
			{
				Type: model.ContentGenerationContentItemTypeImage,
				ImageURL: &model.ImageURL{
					URL: "${YOUR URL HERE}", // Replace with URL
				},
			},
		},
	}

	createResp, err := client.CreateContentGenerationTask(ctx, createReq)
	if err != nil {
		fmt.Printf("create content generation error: %v\n", err)
		return
	}
	fmt.Printf("Task Created with ID: %s\n", createResp.ID)

	fmt.Println("----- get content generation task -----")
	taskID := createResp.ID

	getRequest := model.GetContentGenerationTaskRequest{ID: taskID}

	getResp, err := client.GetContentGenerationTask(ctx, getRequest)
	if err != nil {
		fmt.Printf("get content generation task error: %v\n", err)
		return
	}

	fmt.Printf("Task ID: %s\n", getResp.ID)
	fmt.Printf("Model: %s\n", getResp.Model)
	fmt.Printf("Status: %s\n", getResp.Status)
	fmt.Printf("Failure Reason: %s\n", getResp.FailureReason)
	fmt.Printf("Video URL: %s\n", getResp.Content.VideoURL)
	fmt.Printf("Completion Tokens: %d\n", getResp.Usage.CompletionTokens)
	fmt.Printf("Created At: %d\n", getResp.CreatedAt)
	fmt.Printf("Updated At: %d\n", getResp.UpdatedAt)
}

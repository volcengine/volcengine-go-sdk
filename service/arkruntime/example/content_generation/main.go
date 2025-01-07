package main

import (
	"context"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
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
				Text: volcengine.String("龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原 --ratio 1:1"),
			},
			{
				Type: model.ContentGenerationContentItemTypeImage,
				ImageURL: &model.ImageURL{
					URL: "${YOUR URL HERE}", // Replace with URL
				},
			},
		},
	}

	createResponse, err := client.CreateContentGenerationTask(ctx, createReq)
	if err != nil {
		fmt.Printf("create content generation error: %v\n", err)
		return
	}
	fmt.Printf("Task Created with ID: %s\n", createResponse.ID)

	fmt.Println("----- get content generation task -----")
	taskID := createResponse.ID

	getRequest := model.GetContentGenerationTaskRequest{ID: taskID}

	getResponse, err := client.GetContentGenerationTask(ctx, getRequest)
	if err != nil {
		fmt.Printf("get content generation task error: %v\n", err)
		return
	}

	fmt.Printf("Task ID: %s\n", getResponse.ID)
	fmt.Printf("Model: %s\n", getResponse.Model)
	fmt.Printf("Status: %s\n", getResponse.Status)
	fmt.Printf("Failure Reason: %s\n", getResponse.FailureReason)
	fmt.Printf("Video URL: %s\n", getResponse.Content.VideoURL)
	fmt.Printf("Completion Tokens: %d\n", getResponse.Usage.CompletionTokens)
	fmt.Printf("Created At: %d\n", getResponse.CreatedAt)
	fmt.Printf("Updated At: %d\n", getResponse.UpdatedAt)

	fmt.Println("----- list content generation task -----")

	listRequest := model.ListContentGenerationTasksRequest{
		PageNum:  volcengine.Int(1),
		PageSize: volcengine.Int(10),
		Filter: &model.ListContentGenerationTasksFilter{
			Status: volcengine.String(model.StatusSucceeded),
			// TaskIDs: []string{"cgt-example-1", "cgt-example-2"},
			// Model: "${YOUR_ENDPOINT_ID}",
		},
	}

	listResponse, err := client.ListContentGenerationTasks(ctx, listRequest)
	if err != nil {
		fmt.Printf("failed to list content generation tasks: %v\n", err)
	}

	fmt.Printf("ListContentGenerationTasks returned %v results\n", listResponse.Total)

	fmt.Println("----- delete content generation task -----")

	deleteRequest := model.DeleteContentGenerationTaskRequest{ID: taskID}

	err = client.DeleteContentGenerationTask(ctx, deleteRequest)
	if err != nil {
		fmt.Printf("delete content generation task error: %v\n", err)
	} else {
		fmt.Println("successfully deleted task id: ", taskID)
	}

}

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine"

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
	modelEp := os.Getenv("ENDPOINT_ID")
	//imageURL := os.Getenv("IMAGE_URL")

	fmt.Println("----- create content generation task -----")
	createReq := model.CreateContentGenerationTaskRequest{
		Model: modelEp, // Replace with your endpoint ID
		Content: []*model.CreateContentGenerationContentItem{
			{
				Type: model.ContentGenerationContentItemTypeText,
				Text: volcengine.String("龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原"),
			},
			/*
				{
					Type: model.ContentGenerationContentItemTypeImage,
					ImageURL: &model.ImageURL{
						URL: imageURL, // Replace with URL
					},
					// Role: volcengine.String("first_frame"),
				},
			*/
		},
		// CallbackUrl: volcengine.String("CALLBACK_URL"),
		ServiceTier:           volcengine.String("default"),
		ExecutionExpiresAfter: volcengine.Int64(3600),
	}

	createResponse, err := client.CreateContentGenerationTask(ctx, createReq)
	if err != nil {
		fmt.Printf("create content generation error: %v\n", err)
		return
	}
	fmt.Printf("Task Created with ID: %s\n", createResponse.ID)

	time.Sleep(1 * time.Minute)
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
	fmt.Printf("Video URL: %s\n", getResponse.Content.VideoURL)
	fmt.Printf("Completion Tokens: %d\n", getResponse.Usage.CompletionTokens)
	fmt.Printf("Created At: %d\n", getResponse.CreatedAt)
	fmt.Printf("Updated At: %d\n", getResponse.UpdatedAt)
	if getResponse.Seed != nil {
		fmt.Printf("Seed: %d\n", getResponse.Seed)
	}
	if getResponse.RevisedPrompt != nil {
		fmt.Printf("RevisedPrompt: %s\n", *getResponse.RevisedPrompt)
	}
	if getResponse.ServiceTier != nil {
		fmt.Printf("ServiceTier: %s\n", *getResponse.ServiceTier)
	}
	if getResponse.ExecutionExpiresAfter != nil {
		fmt.Printf("ExecutionExpiresAfter: %d\n", *getResponse.ExecutionExpiresAfter)
	}
	if getResponse.Error != nil {
		fmt.Printf("Error Code: %s\n", getResponse.Error.Code)
		fmt.Printf("Error Message: %s\n", getResponse.Error.Message)
	}

	fmt.Println("----- list content generation task -----")

	listRequest := model.ListContentGenerationTasksRequest{
		PageNum:  volcengine.Int(1),
		PageSize: volcengine.Int(10),
		Filter: &model.ListContentGenerationTasksFilter{
			Status: volcengine.String(model.StatusSucceeded),
			//TaskIDs: volcengine.StringSlice([]string{"cgt-example-1", "cgt-example-2"}),
			//Model:   volcengine.String(modelEp),
			ServiceTier: volcengine.String("default"),
		},
	}

	listResponse, err := client.ListContentGenerationTasks(ctx, listRequest)
	if err != nil {
		fmt.Printf("failed to list content generation tasks: %v\n", err)
	}

	fmt.Printf("ListContentGenerationTasks returned %v results\n", listResponse.Total)

	for _, item := range listResponse.Items {
		if item.ServiceTier != nil {
			fmt.Printf("List Item ServiceTier: %s\n", *item.ServiceTier)
		}
		if item.ExecutionExpiresAfter != nil {
			fmt.Printf("List Item ExecutionExpiresAfter: %d\n", *item.ExecutionExpiresAfter)
		}
	}

	fmt.Println("----- delete content generation task -----")

	deleteRequest := model.DeleteContentGenerationTaskRequest{ID: taskID}

	err = client.DeleteContentGenerationTask(ctx, deleteRequest)
	if err != nil {
		fmt.Printf("delete content generation task error: %v\n", err)
	} else {
		fmt.Println("successfully deleted task id: ", taskID)
	}

	// ---- Flex tier test: create & GET ----
	fmt.Println("----- create content generation task (flex) -----")
	createReqFlex := model.CreateContentGenerationTaskRequest{
		Model: modelEp,
		Content: []*model.CreateContentGenerationContentItem{
			{
				Type: model.ContentGenerationContentItemTypeText,
				Text: volcengine.String("使用 flex 级别进行内容生成测试，验证 service_tier 与 expire 字段"),
			},
		},
		ServiceTier:           volcengine.String("flex"),
		ExecutionExpiresAfter: volcengine.Int64(3600),
	}

	createResponseFlex, err := client.CreateContentGenerationTask(ctx, createReqFlex)
	if err != nil {
		fmt.Printf("create flex content generation error: %v\n", err)
		return
	}
	fmt.Printf("Flex Task Created with ID: %s\n", createResponseFlex.ID)

	fmt.Println("----- get content generation task (flex) -----")
	getFlexReq := model.GetContentGenerationTaskRequest{ID: createResponseFlex.ID}
	getFlexResp, err := client.GetContentGenerationTask(ctx, getFlexReq)
	if err != nil {
		fmt.Printf("get flex content generation task error: %v\n", err)
		return
	}

	if getFlexResp.ServiceTier != nil {
		fmt.Printf("Flex ServiceTier: %s\n", *getFlexResp.ServiceTier)
	}
	if getFlexResp.ExecutionExpiresAfter != nil {
		fmt.Printf("Flex ExecutionExpiresAfter: %d\n", *getFlexResp.ExecutionExpiresAfter)
	}

	// Optional: list flex tasks to observe filtering
	fmt.Println("----- list content generation task (flex) -----")
	listFlexReq := model.ListContentGenerationTasksRequest{
		PageNum:  volcengine.Int(1),
		PageSize: volcengine.Int(10),
		Filter: &model.ListContentGenerationTasksFilter{
			ServiceTier: volcengine.String("flex"),
		},
	}
	listFlexResp, err := client.ListContentGenerationTasks(ctx, listFlexReq)
	if err != nil {
		fmt.Printf("failed to list flex content generation tasks: %v\n", err)
	} else {
		fmt.Printf("List Flex Tasks returned %v results\n", listFlexResp.Total)
		for _, item := range listFlexResp.Items {
			if item.ServiceTier != nil {
				fmt.Printf("Flex List Item ServiceTier: %s\n", *item.ServiceTier)
			}
			if item.ExecutionExpiresAfter != nil {
				fmt.Printf("Flex List Item ExecutionExpiresAfter: %d\n", *item.ExecutionExpiresAfter)
			}
		}
	}

	// cleanup flex task
	fmt.Println("----- delete content generation task (flex) -----")
	delFlexReq := model.DeleteContentGenerationTaskRequest{ID: createResponseFlex.ID}
	if err := client.DeleteContentGenerationTask(ctx, delFlexReq); err != nil {
		fmt.Printf("delete flex content generation task error: %v\n", err)
	} else {
		fmt.Println("successfully deleted flex task id: ", createResponseFlex.ID)
	}

}

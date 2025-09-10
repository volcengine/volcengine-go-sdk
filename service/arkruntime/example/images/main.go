package main

import (
	"context"
	"fmt"
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
 */

func main() {
	client := arkruntime.NewClientWithApiKey("")
	ctx := context.Background()
	imageList := make([]string, 0)

	fmt.Println("----- [Seedream] generate images (response format: url) -----")
	generateReq := model.GenerateImagesRequest{
		Model:          "",
		Image:          imageList,
		Prompt:         "帮我改成修改的更高级一点",
		ResponseFormat: volcengine.String(model.GenerateImagesResponseFormatURL),
		Seed:           volcengine.Int64(1234567890),
		Watermark:      volcengine.Bool(true),
		Size:           volcengine.String("1920x1920"),
		PromptExtend:   volcengine.Bool(true),
	}

	imagesResponse, err := client.GenerateImages(ctx, generateReq)
	if err != nil {
		fmt.Printf("generate images error: %v\n", err)
		return
	}

	if imagesResponse.Error != nil {
		fmt.Printf("Error Code: %s\n", imagesResponse.Error.Code)
		fmt.Printf("Error Message: %s\n", imagesResponse.Error.Message)
	}

	fmt.Printf("Model: %s\n", imagesResponse.Model)
	fmt.Printf("Image URL: %s\n", *imagesResponse.Data[0].Url)
	fmt.Printf("Generated Images: %d\n", imagesResponse.Usage.GeneratedImages)
	fmt.Printf("Created: %d\n", imagesResponse.Created)
}

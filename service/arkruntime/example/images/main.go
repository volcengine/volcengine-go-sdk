package main

import (
	"context"
	"fmt"
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
 */
func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()
	modelEp := "YOUR_ENDPOINT_ID"

	fmt.Println("----- [Seedream] generate images (response format: url) -----")
	generateReq := model.GenerateImagesRequest{
		Model:          modelEp, // Replace with your endpoint ID
		Prompt:         "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原",
		ResponseFormat: volcengine.String(model.GenerateImagesResponseFormatURL),
		Seed:           volcengine.Int64(1234567890),
		Watermark:      volcengine.Bool(true),
		Size:           volcengine.String("512x512"),
		GuidanceScale:  volcengine.Float64(2.5),
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

	fmt.Println("----- [Seedream] generate images (response format: base64) -----")
	generateReq = model.GenerateImagesRequest{
		Model:          modelEp, // Replace with your endpoint ID
		Prompt:         "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原",
		ResponseFormat: volcengine.String(model.GenerateImagesResponseFormatBase64),
		Seed:           volcengine.Int64(1234567890),
		Watermark:      volcengine.Bool(true),
		Size:           volcengine.String("512x512"),
		GuidanceScale:  volcengine.Float64(2.5),
	}

	imagesResponse, err = client.GenerateImages(ctx, generateReq)
	if err != nil {
		fmt.Printf("generate images error: %v\n", err)
		return
	}

	if imagesResponse.Error != nil {
		fmt.Printf("Error Code: %s\n", imagesResponse.Error.Code)
		fmt.Printf("Error Message: %s\n", imagesResponse.Error.Message)
	}

	fmt.Printf("Model: %s\n", imagesResponse.Model)
	fmt.Printf("Image URL: %s\n", *imagesResponse.Data[0].B64Json)
	fmt.Printf("Generated Images: %d\n", imagesResponse.Usage.GeneratedImages)
	fmt.Printf("Created: %d\n", imagesResponse.Created)

	fmt.Println("----- [Seededit] generate images (input format: url) -----")
	generateReq = model.GenerateImagesRequest{
		Model:          modelEp, // Replace with your endpoint ID
		Prompt:         "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原",
		Image:          volcengine.String("YOUR_IMAGE_URL_HERE"), // Replace with your input image URL
		ResponseFormat: volcengine.String(model.GenerateImagesResponseFormatURL),
		Seed:           volcengine.Int64(1234567890),
		Watermark:      volcengine.Bool(true),
		Size:           volcengine.String(model.GenerateImagesSizeAdaptive),
		GuidanceScale:  volcengine.Float64(2.5),
	}

	imagesResponse, err = client.GenerateImages(ctx, generateReq)
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

	fmt.Println("----- [Seededit] generate images (input format: base64) -----")
	generateReq = model.GenerateImagesRequest{
		Model:          modelEp, // Replace with your endpoint ID
		Prompt:         "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原",
		Image:          volcengine.String("YOUR_IMAGE_BASE64_HERE"), // Replace with your input image base64 data url
		ResponseFormat: volcengine.String(model.GenerateImagesResponseFormatURL),
		Seed:           volcengine.Int64(1234567890),
		Watermark:      volcengine.Bool(true),
		Size:           volcengine.String(model.GenerateImagesSizeAdaptive),
		GuidanceScale:  volcengine.Float64(2.5),
	}

	imagesResponse, err = client.GenerateImages(ctx, generateReq)
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

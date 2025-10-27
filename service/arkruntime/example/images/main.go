package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

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
		Model:          modelEp, // Replace with your Seedream endpoint ID
		Prompt:         "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原",
		ResponseFormat: volcengine.String(model.GenerateImagesResponseFormatURL),
		Seed:           volcengine.Int64(1234567890),
		Watermark:      volcengine.Bool(true),
		Size:           volcengine.String("1024x1024"),
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
		Size:           volcengine.String("1024x1024"),
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
		Model:          modelEp, // Replace with your Seededit endpoint ID
		Prompt:         "龙与地下城女骑士背景是起伏的平原，目光从镜头转向平原",
		Image:          volcengine.String("YOUR_IMAGE_BASE64_HERE"), // Replace with your input image base64 data url
		ResponseFormat: volcengine.String(model.GenerateImagesResponseFormatURL),
		Seed:           volcengine.Int64(1234567890),
		Watermark:      volcengine.Bool(true),
		Size:           volcengine.String(model.GenerateImagesSizeAdaptive),
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

	fmt.Println("----- [Seedream] streaming generate images (response format: url) -----")
	var sequentialImageGeneration model.SequentialImageGeneration = "auto"
	maxImages := 9
	generateReq = model.GenerateImagesRequest{
		Model:                     modelEp, // Replace with your Seedream endpoint ID
		Prompt:                    "星球大战, 场面壮观, 需要描述3个连续场面",
		ResponseFormat:            volcengine.String(model.GenerateImagesResponseFormatURL),
		Seed:                      volcengine.Int64(1234567890),
		Watermark:                 volcengine.Bool(true),
		Size:                      volcengine.String("1024x1024"),
		SequentialImageGeneration: &sequentialImageGeneration,
		SequentialImageGenerationOptions: &model.SequentialImageGenerationOptions{
			MaxImages: &maxImages,
		},
	}

	stream, err := client.GenerateImagesStreaming(ctx, generateReq)
	if err != nil {
		fmt.Printf("call GenerateImagesStreaming error: %v\n", err)
		return
	}

	defer stream.Close()

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Stream generate images error: %v\n", err)
			break
		}

		if recv.Type == model.ImageGenerationStreamEventPartialSucceeded {
			fmt.Printf("Stream generate images error: %v\n", recv.Error)
			if strings.EqualFold(recv.Error.Code, "InternalServiceError") {
				break
			}
		}

		if recv.Type == model.ImageGenerationStreamEventPartialSucceeded {
			if recv.Error == nil && recv.Url != nil {
				fmt.Printf("recv.Size: %s, recv.Url: %s\n", recv.Size, *recv.Url)
			}
		}

		if recv.Type == model.ImageGenerationStreamEventCompleted {
			if recv.Error == nil {
				fmt.Printf("recv.Usage: %v\n", *recv.Usage)
			}
		}
	}

	fmt.Println("----- [Seedream] streaming generate images (response format: base64) -----")

	generateReq = model.GenerateImagesRequest{
		Model:                     modelEp, // Replace with your Seedream endpoint ID
		Prompt:                    "星球大战, 场面壮观, 需要描述3个连续场面",
		ResponseFormat:            volcengine.String(model.GenerateImagesResponseFormatBase64),
		Seed:                      volcengine.Int64(1234567890),
		Watermark:                 volcengine.Bool(true),
		Size:                      volcengine.String("1024x1024"),
		SequentialImageGeneration: &sequentialImageGeneration,
		SequentialImageGenerationOptions: &model.SequentialImageGenerationOptions{
			MaxImages: &maxImages,
		},
	}

	stream, err = client.GenerateImagesStreaming(ctx, generateReq)
	if err != nil {
		fmt.Printf("call GenerateImagesStreaming error: %v\n", err)
		return
	}

	defer stream.Close()

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Stream generate images error: %v\n", err)
			break
		}

		if recv.Type == model.ImageGenerationStreamEventPartialSucceeded {
			fmt.Printf("Stream generate images error: %v\n", recv.Error)
			if strings.EqualFold(recv.Error.Code, "InternalServiceError") {
				break
			}
		}

		if recv.Type == model.ImageGenerationStreamEventPartialSucceeded {
			if recv.Error == nil && recv.B64Json != nil {
				fmt.Printf("recv.Size: %s, recv.B64Json: %s\n", recv.Size, *recv.B64Json)
			}
		}

		if recv.Type == model.ImageGenerationStreamEventCompleted {
			if recv.Error == nil {
				fmt.Printf("recv.Usage: %v\n", *recv.Usage)
			}
		}
	}
}

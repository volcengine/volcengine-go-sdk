package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

func main() {
	client := arkruntime.NewClientWithApiKey(
		os.Getenv("ARK_API_KEY"),
	)
	ctx := context.Background()

	fmt.Println("----- multimodal embeddings request -----")
	req := model.MultiModalEmbeddingRequest{
		Model: "${YOUR_ENDPOINT_ID}",
		Input: []model.MultimodalEmbeddingInput{
			{
				Type:     model.MultiModalEmbeddingInputTypeImageURL,
				ImageURL: &model.MultimodalEmbeddingImageURL{URL: "https://ark-project.tos-cn-beijing.volces.com/images/view.jpeg"},
			},
		},
	}

	resp, err := client.CreateMultiModalEmbeddings(ctx, req)
	if err != nil {
		fmt.Printf("multimodal embeddings error: %v\n", err)
		return
	}

	s, _ := json.Marshal(resp)
	fmt.Println(string(s))
}

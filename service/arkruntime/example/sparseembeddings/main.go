package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func main() {
	client := arkruntime.NewClientWithApiKey(
		os.Getenv("ARK_API_KEY"),
	)
	ctx := context.Background()

	fmt.Println("----- multimodal embeddings request -----")
	req := model.MultiModalEmbeddingRequest{
		Model: "doubao-embedding-vision-250615",
		Input: []model.MultimodalEmbeddingInput{
			{
				Type: model.MultiModalEmbeddingInputTypeText,
				Text: volcengine.String("花椰菜又称菜花、花菜，是一种常见的蔬菜。"),
			},
		},
		SparseEmbedding: &model.SparseEmbeddingInput{
			Type: model.SparseEmbeddingInputTypeEnabled,
		},
	}

	resp, err := client.CreateMultiModalEmbeddings(ctx, req)
	if err != nil {
		fmt.Printf("multimodal embeddings error: %v\n", err)
		return
	}

	s, _ := json.Marshal(resp.Data.SparseEmbedding)
	fmt.Println(string(s))
}

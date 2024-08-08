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

	fmt.Println("----- tokenization request -----")
	req := model.TokenizationRequestStrings{
		Text: []string{
			"花椰菜又称菜花、花菜，是一种常见的蔬菜。",
		},
		Model: "${YOUR_ENDPOINT_ID}",
	}

	resp, err := client.CreateTokenization(ctx, req)
	if err != nil {
		fmt.Printf("tokenization error: %v\n", err)
		return
	}

	s, _ := json.Marshal(resp)
	fmt.Println(string(s))
}

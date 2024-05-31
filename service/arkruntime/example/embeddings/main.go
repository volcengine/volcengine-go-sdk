package main

import (
	"context"
	"encoding/json"
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

	fmt.Println("----- embeddings request -----")
	req := model.EmbeddingRequestStrings{
		Input: []string{
			"The food was delicious and the waiter",
			"Other examples of embedding request",
		},
		Model: "${YOUR_ENDPOINT_ID}",
	}

	resp, err := client.CreateEmbeddings(ctx, req)
	if err != nil {
		fmt.Printf("embeddings error: %v\n", err)
		return
	}

	s, _ := json.Marshal(resp)
	fmt.Println(string(s))

}

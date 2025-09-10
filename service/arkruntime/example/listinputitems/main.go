package main

import (
	"context"
	"fmt"
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

const (
	defaultLimit = 10
)

/**
 * Authentication
 * If you authorize your endpoint using an API key, you can set your api key to environment variable "ARK_API_KEY"
 * client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
 * Note: If you use an API key, this API key will not be refreshed.
 * To prevent the API from expiring and failing after some time, choose an API key with no expiration date.
 */

func main() {
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
	ctx := context.Background()

	responseID := "{RESPONSE_ID}"
	// list input items
	listResponsesReq := &responses.ListInputItemsRequest{
		Limit: volcengine.Int32(defaultLimit),
	}

	resp, err := client.ListResponseInputItems(ctx, responseID, listResponsesReq)
	if err != nil {
		fmt.Printf("list error: %v\n", err)
		return
	}
	for _, item := range resp.Data {
		fmt.Println(item)
	}

	// do not have any more data
	if !resp.GetHasMore() {
		return
	}

	// get next page data
	id := resp.GetFirstId()
	listResponsesReq.Before = volcengine.String(id)

	resp, err = client.ListResponseInputItems(ctx, responseID, listResponsesReq)
	if err != nil {
		fmt.Printf("list error: %v\n", err)
		return
	}
	for _, item := range resp.Data {
		fmt.Println(item)
	}
}

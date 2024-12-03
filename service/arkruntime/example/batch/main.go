package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func main() {
	timeout := time.Minute * 30
	workerNum := 10000

	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"), arkruntime.WithTimeout(timeout))
	client.StartBatchWorker(workerNum)

	ctx := context.Background()
	i := 0
	taskNum := 5
	wg := sync.WaitGroup{}
	for i < workerNum {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			j := 0
			for j < taskNum {
				_, err := client.CreateBatchChatCompletion(ctx, model.ChatCompletionRequest{
					Model: os.Getenv("ENDPOINT_ID"),
					Messages: []*model.ChatCompletionMessage{
						{
							Role: model.ChatMessageRoleSystem,
							Content: &model.ChatCompletionMessageContent{
								StringValue: volcengine.String("你是豆包，是由字节跳动开发的 AI 人工智能助手"),
							},
						},
						{
							Role: model.ChatMessageRoleUser,
							Content: &model.ChatCompletionMessageContent{
								StringValue: volcengine.String("常见的十字花科植物有哪些？"),
							},
						},
					},
				})
				j++
				if err != nil {
					fmt.Printf("worker %d request %d Fail Err %s\n", index, j, err)
					continue
				}
			}
		}(i)
		i++
	}
	wg.Wait()
}

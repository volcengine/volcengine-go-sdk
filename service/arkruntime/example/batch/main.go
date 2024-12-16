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
	timeout := time.Hour
	workerNum := 10000

	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"), arkruntime.WithTimeout(timeout))
	client.StartBatchWorker(workerNum)

	ctx := context.Background()
	taskNum := 5
	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for j := 0; j < taskNum; j++ {
				resp, err := client.CreateBatchChatCompletion(ctx, model.ChatCompletionRequest{
					Model: "${YOUR_BOT_ID}",
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
				if err != nil {
					fmt.Printf("worker %d request %d Fail Err %s\n", index, j, err)
					continue
				}
				fmt.Println(*resp.Choices[0].Message.Content.StringValue)
			}
		}(i)
	}
	wg.Wait()
	client.StopBatchWorker()
}

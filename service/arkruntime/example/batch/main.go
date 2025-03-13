package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func main() {
	// 使用 API 密钥创建一个新的客户端实例，并设置超时时间
	client := arkruntime.NewClientWithApiKey(
		"YOUR_API_KEY",
		arkruntime.WithBatchMaxParallel(3000), // 设置发起请求的最大并发数量为 3000
	)

	// 这里手动构造 50000 个请求，在实际应用中，你可以从文件、消息队列或数据库中加载真实的请求
	requests := MockRequests("YOUR_ENDPOINT_ID", 50000)

	wg, ctx := sync.WaitGroup{}, context.Background()

	// 可以在这里设置全局的超时时间，如果超过这个时间，所有的请求都会被取消
	ctx, cancel := context.WithTimeout(ctx, 24*time.Hour)
	defer cancel()

	// 发起请求
	for request := range requests {
		wg.Add(1)
		// 异步发起请求
		go func(request model.CreateChatCompletionRequest) {
			defer wg.Done()

			// 可以在这里设置每个请求的超时时间，如果超过这个时间，这个请求会被取消
			ctx, cancel := context.WithTimeout(ctx, 24*time.Hour)
			defer cancel()

			// 发起批量推理请求
			result, err := client.CreateBatchChatCompletion(ctx, request)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Println(MustMarshalJson(result))
			}
		}(request)
	}
	// 等待所有协程完成任务
	wg.Wait()
}

// MockRequests 模拟生成请求，这里只是简单地生成了 count 个相同的请求。
// 在实际应用中，你可以从文件、消息队列或数据库中加载真实的请求。
func MockRequests(endpoint string, count int) <-chan model.CreateChatCompletionRequest {
	requests := make(chan model.CreateChatCompletionRequest)

	go func() {
		defer close(requests)
		for i := 0; i < count; i++ {
			requests <- model.CreateChatCompletionRequest{
				Model: endpoint,
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
			}
		}
	}()

	return requests
}

func MustMarshalJson(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

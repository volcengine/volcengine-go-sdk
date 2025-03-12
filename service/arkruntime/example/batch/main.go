package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func main() {
	// 设置超时时间为 24 小时，以利用潮汐资源
	ctx, cancel := context.WithTimeout(context.Background(), 24*time.Hour)
	defer cancel()

	// 使用 API 密钥创建一个新的客户端实例
	client := arkruntime.NewClientWithApiKey(
		os.Getenv("ARK_API_KEY"),
		arkruntime.WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				// 设置发起请求的最大并发数量为 3000
				MaxConnsPerHost: 3000,
			},
		}),
	)

	wg := sync.WaitGroup{}

	// 发起 50000 个请求
	for i := 0; i < 50000; i++ {
		wg.Add(1)

		// 异步发起请求
		go func(index int) {
			defer wg.Done()

			// 发起批量推理请求
			result, err := client.CreateBatchChatCompletion(ctx, model.ChatCompletionRequest{
				Model: os.Getenv("YOUR_ENDPOINT_ID"),
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
				fmt.Fprintln(os.Stderr, index, err)
			} else {
				fmt.Println(index, MustMarshalJson(result))
			}
		}(i)
	}
	// 等待所有工作线程完成任务
	wg.Wait()
}

func MustMarshalJson(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

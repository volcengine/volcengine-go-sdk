package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//os.Setenv("ARK_API_KEY", "e0b31760-c725-4478-8a7a-4126df8c37b2")
	//os.Setenv("ARK_API_KEY", "171d8284-ca9d-4da4-989a-fa9cd3bb1216")
	retryTimes, err := strconv.Atoi(os.Getenv("RETRY_TIMES"))
	if err != nil {
		retryTimes = 5
	}
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = 10 * time.Minute
	}

	workerNum, err := strconv.Atoi(os.Getenv("WORKER_NUM"))
	if err != nil {
		workerNum = 100
	}
	taskNum, err := strconv.Atoi(os.Getenv("TASK_NUM"))
	if err != nil {
		taskNum = 5
	}
	client := arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"), arkruntime.WithRetryTimes(retryTimes), arkruntime.WithBaseUrl("https://ark-stg.cn-beijing.volces.com/api/v3"), arkruntime.WithTimeout(timeout), arkruntime.WithHTTPClient(&http.Client{Timeout: timeout, Transport: &http.Transport{
		IdleConnTimeout: timeout,
		MaxIdleConns:    workerNum,
		MaxConnsPerHost: workerNum,
	}}))

	client.StartBatchWorker(workerNum)

	ctx := context.Background()
	i := 0
	wg := sync.WaitGroup{}
	start := time.Now()
	successCount := atomic.Int32{}
	failCount := atomic.Int32{}
	for i < workerNum {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			j := 0
			fmt.Printf("worker %d start\n", index)
			for j < taskNum {
				_, err := client.CreateBatchChatCompletion(ctx, model.ChatCompletionRequest{
					//Model: "ep-20241121212908-lv8jf",
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
					failCount.Add(1)
					continue
				}
				successCount.Add(1)
				fmt.Printf("worker %d request %d Success\n", index, j)
			}
		}(i)
		i++
	}
	wg.Wait()
	fmt.Printf("耗时: %fs\n", time.Now().Sub(start).Seconds())
	fmt.Printf("成功任务数: %d\n", successCount.Load())
	fmt.Printf("失败任务数: %d\n", failCount.Load())
}

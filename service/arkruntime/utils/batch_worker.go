package utils

import (
	"context"
	"math"
	"math/rand"
	"sync"
	"time"
)

type BatchTask struct {
	Model      string
	DoneChan   chan error
	TaskFunc   func() (BatchTaskResult, error)
	RetryTimes int
	Context    context.Context
}

type BatchTaskResult struct {
	NeedRetry    bool
	RequeueAfter int64
}

type BatchWorkerPool struct {
	workerNum    int
	taskQueue    chan BatchTask
	breakerLock  sync.Mutex
	modelBreaker sync.Map
	retryPolicy  RetryPolicy
}

func NewBatchWorkerPool(workerNum int) *BatchWorkerPool {
	return &BatchWorkerPool{
		workerNum:    workerNum,
		taskQueue:    make(chan BatchTask, workerNum),
		breakerLock:  sync.Mutex{},
		modelBreaker: sync.Map{},
		retryPolicy: RetryPolicy{
			InitialBackoff: 500 * time.Millisecond,
			MaxBackoff:     5 * time.Second,
		},
	}
}

func (p *BatchWorkerPool) Submit(ctx context.Context, model string, taskFunc func() (BatchTaskResult, error), doneChan chan error) {
	p.taskQueue <- BatchTask{
		Context:  ctx,
		Model:    model,
		DoneChan: doneChan,
		TaskFunc: taskFunc,
	}
}

func (p *BatchWorkerPool) Run() {
	wg := sync.WaitGroup{}
	for i := 0; i < p.workerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				task, ok := <-p.taskQueue
				if !ok {
					break
				}
				select {
				// 任务被取消
				case <-task.Context.Done():
					task.DoneChan <- task.Context.Err()
				default:
					breaker := p.GetOrCreateBreaker(task.Model)
					if !breaker.IsAllowed() {
						p.AddAfter(task, breaker.GetAllowedDuration())
						continue
					}
					result, err := task.TaskFunc()
					if !result.NeedRetry {
						task.DoneChan <- err
						continue
					}
					task.RetryTimes++
					if result.RequeueAfter > 0 {
						p.AddAfter(task, time.Duration(result.RequeueAfter)*time.Second)
						breaker.Reset(time.Duration(result.RequeueAfter) * time.Second)
						continue
					}
					p.AddAfter(task, p.getRetryDuration(task.RetryTimes))
				}
			}
		}()
	}
	wg.Wait()
}

func (p *BatchWorkerPool) getRetryDuration(retryTimes int) time.Duration {
	sleepSeconds := math.Min(p.retryPolicy.InitialBackoff.Seconds()*math.Pow(2.0, float64(retryTimes)), p.retryPolicy.MaxBackoff.Seconds())
	jitter := 1.0 - 0.25*rand.Float64()
	return time.Duration(sleepSeconds*jitter) * time.Second
}

func (p *BatchWorkerPool) AddAfter(task BatchTask, duration time.Duration) {
	time.AfterFunc(duration, func() {
		p.taskQueue <- task
	})
}

func (p *BatchWorkerPool) Close() {
	close(p.taskQueue)
}

func (p *BatchWorkerPool) GetOrCreateBreaker(model string) *Breaker {
	if value, ok := p.modelBreaker.Load(model); ok {
		if breaker, ok := value.(*Breaker); ok {
			return breaker
		}
	}
	breaker := NewBreaker()
	p.modelBreaker.Store(model, breaker)
	return breaker
}

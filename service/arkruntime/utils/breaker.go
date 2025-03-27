package utils

import (
	"math"
	"sync"
	"time"

	"github.com/google/uuid"
)

type querySet struct {
	index map[uuid.UUID]int
	items []uuid.UUID
	lock  sync.Mutex
}

// newQuerySet 创建一个新的 QuerySet 实例
func newQuerySet() *querySet {
	return &querySet{
		index: make(map[uuid.UUID]int),
		items: make([]uuid.UUID, 0),
	}
}

// Query 查询元素的索引
func (qs *querySet) Query(item uuid.UUID) int {
	qs.lock.Lock()
	defer qs.lock.Unlock()
	return qs.index[item]
}

// Add 向 QuerySet 中添加元素
func (qs *querySet) Add(item uuid.UUID) {
	qs.lock.Lock()
	defer qs.lock.Unlock()

	if _, ok := qs.index[item]; ok {
		return
	}
	qs.items = append(qs.items, item)
	qs.index[item] = len(qs.items) - 1
}

// Remove 从 QuerySet 中移除元素
func (qs *querySet) Remove(item uuid.UUID) {
	qs.lock.Lock()
	defer qs.lock.Unlock()

	index, ok := qs.index[item]
	if !ok {
		return
	}

	qs.items[index] = qs.items[len(qs.items)-1]
	qs.index[qs.items[index]] = index
	qs.items = qs.items[:len(qs.items)-1]
	delete(qs.index, item)
}

// Breaker 定义熔断结构体
type Breaker struct {
	allowTime time.Time
	waiters   *querySet
	lock      sync.Mutex
}

// NewBreaker 创建一个新的 Breaker 实例
func NewBreaker() *Breaker {
	return &Breaker{
		allowTime: time.Now(),
		waiters:   newQuerySet(),
	}
}

// allow 判断是否允许通过
// 慢启动策略：
// 1. 初始允许通过的请求数量为 1；
// 2. 每经过 1 秒，允许通过的请求数量增加一倍；
// 3. 10s 后，允许所有请求通过。
// 由于并发的随机性，实际允许通过的请求数量可能会超过预期。
// 但在并发度较高时，慢启动策略可以有效地减少初始阶段并发请求的数量。
func (mb *Breaker) allow(id uuid.UUID) bool {
	mb.lock.Lock()
	defer mb.lock.Unlock()
	cur := time.Now()
	elapsed := cur.Sub(mb.allowTime).Seconds()

	if elapsed <= 0 {
		return false
	}
	if elapsed > 10 {
		return true
	}

	return float64(mb.waiters.Query(id)) < math.Pow(2, elapsed)
}

// getAllowedDuration 计算允许的持续时间
func (mb *Breaker) getAllowedDuration() time.Duration {
	mb.lock.Lock()
	defer mb.lock.Unlock()
	allowDuration := time.Until(mb.allowTime)
	if allowDuration < time.Second {
		return time.Second
	}
	return allowDuration
}

// Reset 重置 allowTime
func (mb *Breaker) Reset(duration time.Duration) {
	mb.lock.Lock()
	defer mb.lock.Unlock()
	mb.allowTime = time.Now().Add(duration)
}

func (mb *Breaker) acquire() uuid.UUID {
	id := uuid.New()
	mb.waiters.Add(id)
	return id
}

func (mb *Breaker) release(id uuid.UUID) {
	mb.waiters.Remove(id)
}

// Wait 同步等待直到允许通过
func (mb *Breaker) Wait() {
	id := mb.acquire()
	defer mb.release(id)

	for !mb.allow(id) {
		time.Sleep(mb.getAllowedDuration())
	}
}

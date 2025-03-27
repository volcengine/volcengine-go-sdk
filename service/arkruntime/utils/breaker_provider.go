package utils

import (
	"sync"
)

type ModelBreakerProvider struct {
	breakers sync.Map
	lock     sync.Mutex
}

func NewModelBreakerProvider() *ModelBreakerProvider {
	return &ModelBreakerProvider{
		breakers: sync.Map{},
	}
}

func (p *ModelBreakerProvider) GetOrCreateBreaker(model string) *Breaker {
	if value, ok := p.breakers.Load(model); ok {
		if breaker, ok := value.(*Breaker); ok {
			return breaker
		}
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	if value, ok := p.breakers.Load(model); ok {
		if breaker, ok := value.(*Breaker); ok {
			return breaker
		}
	}

	breaker := NewBreaker()
	p.breakers.Store(model, breaker)
	return breaker
}

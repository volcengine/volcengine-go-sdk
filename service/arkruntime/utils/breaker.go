package utils

import "time"

type Breaker struct {
	LastTime time.Time
}

func NewBreaker() *Breaker {
	return &Breaker{
		LastTime: time.Now().Add(-time.Second),
	}
}

func (b *Breaker) IsAllowed() bool {
	return time.Now().After(b.LastTime)
}

func (b *Breaker) Reset(Duration time.Duration) {
	b.LastTime = time.Now().Add(Duration)
}

func (b *Breaker) GetAllowedDuration() time.Duration {
	return time.Until(b.LastTime)
}

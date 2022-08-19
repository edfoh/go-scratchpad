package ratelimit

import (
	"sync"
	"time"
)

var now = time.Now

type RateLimitingStrategy interface {
	TakeN(n int, availableCredits int) bool
}

type TokenBucket struct {
	sync.Mutex

	burstCapacity       int
	capacity            int
	current             int
	lastUpdated         time.Time
	refillRatePerSecond float64
}

func NewTokenBucket(capacity int, refillRatePerSecond float64) RateLimitingStrategy {
	return &TokenBucket{
		capacity:            capacity,
		refillRatePerSecond: refillRatePerSecond,
		current:             capacity,
		lastUpdated:         now(),
	}
}

func (b *TokenBucket) TakeN(n int, availableCredits int) bool {
	b.Lock()
	defer b.Unlock()

	b.refill()

	if b.current-n >= 0 {
		b.current -= n
		return true
	}
	return false
}

func (b *TokenBucket) refill(availableCredits int) {

	timeElapsed := now().Sub(b.lastUpdated)

	refillTokens := timeElapsed.Seconds() * b.refillRatePerSecond

	temp := max(b.current+int(refillTokens), b.capacity)
	temp = max(temp+availableCredits, b.burstCapacity)

	b.current = temp
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

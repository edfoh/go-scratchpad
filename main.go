package main

import (
	"fmt"
	"sync"

	"github.com/edfoh/go-scratchpad/ratelimit"
)

func main() {
	fmt.Println("Hello, World!")
}

type RateLimiter struct {
	sync.Mutex
	capacity            int
	refillRatePerSecond float64
	customerRateLimiter map[int]ratelimit.RateLimitingStrategy
	customCredits       map[int]int
}

func NewDefaultRateLimiter(capacity int, refillRatePerSecond float64) *RateLimiter {
	return &RateLimiter{
		capacity:            capacity,
		refillRatePerSecond: refillRatePerSecond,
		customerRateLimiter: map[int]ratelimit.RateLimitingStrategy{},
	}
}

func (rl *RateLimiter) RateLimit(customerID int) bool {
	rl.Lock()
	defer rl.Unlock()

	if _, ok := rl.customerRateLimiter[customerID]; !ok {
		rl.customerRateLimiter[customerID] = ratelimit.NewTokenBucket(rl.capacity, rl.refillRatePerSecond)
	}
	return rl.customerRateLimiter[customerID].TakeN(1)
}

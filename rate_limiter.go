package main

import (
	"fmt"
	"sync"
	"time"
)

type UserRateLimiter struct {
	mutex            sync.Mutex
	maxCapacity      float64
	tokensAvailable  float64
	lastRefilledTime time.Time
	refilledRate     float64
}

func NewUserRateLimiter(maxCapacity, refilledRate float64) *UserRateLimiter {
	return &UserRateLimiter{
		maxCapacity:      maxCapacity,
		refilledRate:     refilledRate,
		lastRefilledTime: time.Now(),
		tokensAvailable:  maxCapacity,
	}
}

func (ul *UserRateLimiter) IsRequestAllowed() bool {
	ul.mutex.Lock()
	defer ul.mutex.Unlock()

	elapsedTime := time.Since(ul.lastRefilledTime)
	if elapsedTime.Seconds() > 0 {
		fmt.Printf("Attempt refilling the rate bucket, Duration : %f, Rate : %f\n",
			elapsedTime.Seconds(), ul.refilledRate)
		refilledCapacity := elapsedTime.Seconds() * ul.refilledRate
		fmt.Printf("Refilling capacity : %f\n", refilledCapacity)
		ul.tokensAvailable += refilledCapacity
		if ul.tokensAvailable > ul.maxCapacity {
			ul.tokensAvailable = ul.maxCapacity
		}

		ul.lastRefilledTime = time.Now()
	}

	fmt.Println("Tokens Available :", ul.tokensAvailable)
	if ul.tokensAvailable < 1 {
		fmt.Println("No tokens available. Request not allowed")
		return false
	}

	ul.tokensAvailable--

	return true
}

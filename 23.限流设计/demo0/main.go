package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	// rate.NewLimiter(100, 1) // 桶的大小为1, 每秒向桶里放100个token
	limit := rate.Every(100 * time.Microsecond) // 每隔100ms放一个token, 相当于1秒放10个token
	limiter := rate.NewLimiter(limit, 1)
	limiter.Wait(context.Background())
	limiter.Allow()
	rev := limiter.Reserve()
	if !rev.OK() {
		rev.Delay()
	}
	limiter.SetLimit(100)
	limiter.SetBurst(100)
}

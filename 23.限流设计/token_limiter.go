package _3_限流设计

import "time"

type TokenLimiterService struct {
	token chan struct{}
}
func NewTokenLimiterService(interval time.Duration, maxCount int) *TokenLimiterService {
	t := TokenLimiterService{
		token:make(chan struct{}, maxCount),
	}
	go func() {
		ticker := time.NewTicker(time.Nanosecond * maxCount/ interval * 1000 * 1000)
		for {
			<-ticker.C
			t.token <- struct{}{}
		}
	}()
	return &t
}

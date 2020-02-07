package _3_限流设计

import (
	"fmt"
	"testing"
	"time"
)

func TokenBucket(limit, rate int) chan struct{} {
	bucket := make(chan struct{}, limit)
	ticker := time.NewTicker(time.Second * time.Duration(1))
	for i := 0; i < limit; i ++ {
		bucket <- struct{}{}
	}
	go func() {
		for {
			for i := 0; i < rate; i ++ {
				bucket <- struct{}{}
			}
			fmt.Println("向桶放入:", rate)
			<-ticker.C
		}
	}()
	return bucket
}
func PopBucket(bucket <-chan struct{}, n int) {
	for i := 0; i < n; i ++ {
		select {
		case s, ok := <-bucket:
			if ok {
				fmt.Println("消费", s)
			}
		}
		//s, ok := <- bucket
		//if ok {
		//	fmt.Println("消费", s)
		//}
	}
}
func TestTokenBucket(t *testing.T) {
	rate := 10
	bucket := TokenBucket(20, rate)
	dateLen := 100
	for i := 0; i <= dateLen; i ++ {
		PopBucket(bucket, i)
	}
}


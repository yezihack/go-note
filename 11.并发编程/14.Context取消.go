package main

import (
	"context"
	"fmt"
	"time"
)

func isCancelContext(ctx context.Context) bool {
	select {
	case <-ctx.Done()://接受取消通知
		return true
	default:
		return false
	}
}

func main() {
	//创建一个上下文,父上下文
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i ++ {
		go func(i int) {
			for {
				if isCancelContext(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 10)
			}
			fmt.Println("context,i=", i)
		}(i)
	}
	cancel()
	time.Sleep(time.Millisecond *100)
}
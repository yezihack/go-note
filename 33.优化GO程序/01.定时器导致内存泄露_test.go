package _3_优化GO程序

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)
//设置
func ProcessChannelMsg(ctx context.Context, in <-chan string) {
	duration := time.Second * 1
	delay := time.NewTimer(duration)
	defer delay.Stop()
	for {
		delay.Reset(duration)
		select {
		case s, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(s)
		case <-delay.C:
			fmt.Println("超时啦")
		case <-ctx.Done():
			fmt.Println("ctx Done")
			return
		}
	}

}
func TestProcessChannelMsg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	go func() {
		for {
			fmt.Println("goroutine Num: ", runtime.NumGoroutine())
			time.Sleep(time.Millisecond * 500)
		}
	}()
	ch := make(chan string, 1)
	ch <- "ok"
	ProcessChannelMsg(ctx, ch)
}
//错误的使用超时设置
//After本质是time.NewTimer时间是不会回收的
//被遗弃的time.After定时任务还是在时间堆里面，定时任务未到期之前，是不会被gc清理的。
func ProcessChannel(ctx context.Context, ch <-chan string) {
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(s)
		case <-time.After(time.Second):
			fmt.Println("超时啦")
		case <-ctx.Done():
			fmt.Println("退出啦")
			return
		}
	}

}
func TestProcessChannel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
	defer cancel()
	go func() {
		for {
			fmt.Println("goroutine Num: ", runtime.NumGoroutine())
			mem := runtime.MemStats{}
			fmt.Printf("%#v", mem)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	ch := make(chan string, 1)
	ch <- "japan"
	ProcessChannel(ctx, ch)
}

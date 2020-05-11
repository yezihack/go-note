package _9_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
Context 内置两个基础的
func Background() Context {
	return background
}

func TODO() Context {
	return todo
}

background一般用于main, 或初使化
todo其它场景新建context使用

参考:https://bettertxt.top/post/golang-concurrent/
*/
func TestWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Finish over")
				return
			default:
				fmt.Println("Doing")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("start quit")
	cancel()
}
func TestWithMoreContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go watch(ctx, i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("finish over")
			return
		default:
			fmt.Println("Doing", i)
			time.Sleep(1 * time.Second)
		}
	}
}

// withDeadline 设置截止时间.
func TestContextTimeout(t *testing.T) {
	//over := time.Now().Add(2 * time.Second)
	over := time.Now().Add(4 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), over)
	defer cancel()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("over")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// withDeadline 设置截止时间.
func TestContextTimeout2(t *testing.T) {
	//over := time.Now().Add(2 * time.Second)
	over := time.Now().Add(50 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), over)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("over")
			default:
				fmt.Println("doing")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 10)

	fmt.Println("game over")
	fmt.Println("game over")
}

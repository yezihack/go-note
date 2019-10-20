package demo4

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	var ch = make(chan int)
	go func() {
		rand.Seed(time.Now().UnixNano())
		ch <- rand.Intn(100)
	}()
	fmt.Println(<-ch) //一直会等待ch 里有值,会执行完成. 如果ch里没有值,会一直阻塞住

}

func TestChan2(t *testing.T) {
	ch := make(chan struct{})
	<-ch
}

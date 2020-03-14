package _2_atomic

import (
	"fmt"
	"testing"
	"time"
)

func TestNewEtcdClient(t *testing.T) {
	for i := 0;i < 10; i ++ {
		go func(i int) {
			fmt.Printf("%p\n", Client())
		}(i)
	}
	time.Sleep(time.Second)
}

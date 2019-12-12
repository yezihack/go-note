package _3_限流设计

import (
	"testing"
	"time"
)

func TestNewCountLimitService(t *testing.T) {
	c := NewCountLimitService(time.Millisecond * 500, 100)

	for i := 0; i < 1000; i ++ {
		if !c.GetToken() {
			t.Fatal("拒绝服务", i)
		} else {
			t.Log("通过")
		}
	}
}

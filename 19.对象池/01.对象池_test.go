package _9_对象池

import (
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	pool := NewPool(5)
	for i := 0;i < 10; i ++ {
		v, err := pool.GetObj(time.Second * 1)
		if err != nil {
			t.Fatal(err)
		} else {
			//还回资源
			if err := pool.ReleaseObj(v); err != nil {
				t.Fatal(err)
			}
		}
	}
}

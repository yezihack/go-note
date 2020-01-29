package main

import (
	"time"
	"unsafe"
)

type data struct {
	x [1024*100]byte
}
func dataTest() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}
//go build -o test pointer.go && GODEBUG="gctrace=1" ./test
func main() {
	const N = 1000
	cache := new([N]uintptr)
	for i := 0; i < N; i ++ {
		cache[i] = dataTest()
		time.Sleep(time.Millisecond)
	}
}

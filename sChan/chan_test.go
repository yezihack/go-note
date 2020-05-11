package sChan

import (
	"bytes"
	"runtime"
	"testing"
)

func TestStudyChan(t *testing.T) {
	StudyChan()
}

func TestGoWait(t *testing.T) {
	GoWait()
}
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestGoCount(t *testing.T) {
	n := runtime.GOMAXPROCS(0)
	GoCount(n)
}
func TestCreateCount(t *testing.T) {
	n := runtime.GOMAXPROCS(0)
	CreateCount(n)
}

func TestGoSched(t *testing.T) {
	GoSched()
}

func TestSignChan(t *testing.T) {
	SignChan()
}

func TestCacheChan(t *testing.T) {
	CacheChan()
}

func TestForChan(t *testing.T) {
	ForChan()
}

func TestRandChan(t *testing.T) {
	RandChan()
}

func TestRandDefaultChan(t *testing.T) {
	RandDefaultChan()
}
func TestShowReceive(t *testing.T) {
	for i := 0; i < 10; i++ {
		ShowReceive(i)
	}
}
func TestShowNewPool(t *testing.T) {
	ShowNewPool()
}
func BenchmarkNewPool(b *testing.B) {
	p := NewPool(10)
	for i := 0; i < b.N; i++ {
		p.Put([]byte("a"))
	}
}

func BenchmarkString(b *testing.B) {
	var bb bytes.Buffer
	for i := 0; i < b.N; i++ {
		bb.Write([]byte("a"))
	}
}

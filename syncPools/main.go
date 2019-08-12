package main

import (
	"sync"
	"fmt"
)

var pool = sync.Pool{
	New: func() interface{} {
		fmt.Println("pool")
		b := make([]byte, 8)
		return &b
	},
}

func main() {
	pool.Put(pool.Get().(*[]byte))
	pool.Put(pool.Get().(*[]byte))
	pool.Put(pool.Get().(*[]byte))
	pool.Put(pool.Get().(*[]byte))
	pool.Put(pool.Get().(*[]byte))
	pool.Put(pool.Get().(*[]byte))
	b := pool.Get().(*[]byte)
	fmt.Println(&b)
	pool.Put(b)
	b1 := pool.Get().(*[]byte)
	pool.Put(b1)
	fmt.Println(&b1)
	b2 := pool.Get().(*[]byte)
	pool.Put(b2)
	fmt.Println(&b2)
}
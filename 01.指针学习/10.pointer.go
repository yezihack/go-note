package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"sync/atomic"
	"time"
	"unsafe"
)

type ClientV3 struct {
	name string
	pass string
}

var clientPointer unsafe.Pointer

func Client() (*ClientV3, error) {
	for {
		p := (*ClientV3)(atomic.LoadPointer(&clientPointer))
		if p != nil {
			return p, nil
		}
		client := new(ClientV3)
		client.name = "root"
		client.pass = "root"
		if !atomic.CompareAndSwapPointer(&clientPointer, nil, unsafe.Pointer(client)) {
			time.Sleep(time.Nanosecond * 5)
			continue
		}
		return client, nil
	}
}
func main() {
	c, err := Client()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.name, c.pass)
}

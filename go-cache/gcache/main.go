package main

import (
	"github.com/bluele/gcache"
	"fmt"
)

func main() {
	gc := gcache.New(20).LFU().Build()
	gc.Set("gcache", "this is gcache")
	value, err := gc.Get("gcache")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

}
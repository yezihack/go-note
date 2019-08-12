package main

import (
	"fmt"
	"github.com/ThreeKing2018/goCache"
	"sync"
	"time"
)

type Man struct {
	Cache goCache.GoCacher
	lock  sync.RWMutex
	once  sync.Once
}

var m Man

func getMan() *Man {
	var m Man
	m = Man{}
	m.once.Do(func() {
		m.Cache = goCache.NewDefault()
	})
	return &m
}

func main() {
	getMan().Cache.Set("aa", 100, time.Second*10)
	fmt.Println(getMan().Cache.Get("aa"))
}

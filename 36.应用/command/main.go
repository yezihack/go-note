package main

import (
	"fmt"
	"sync"
)

//入口
func main() {
	g := NewGen()
	g.Show()
}

type Gen struct {
	Mux  sync.RWMutex
	Name string
}

func NewGen() *Gen {
	return &Gen{}
}

// 显示出来
func (g *Gen) Show() {
	fmt.Println("gen")
}

// 设置Set
func (g *Gen) Set() {
	fmt.Println("Set")
}

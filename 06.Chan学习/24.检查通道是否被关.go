package main

import (
	"runtime/debug"
)

type TT int
func IsClose(c chan TT) bool{
	select {
	case <-c:
		return true
	default:
	}
	return false
}
func SaveIsClose(c chan TT, value TT) bool {
	send := false
	defer func() {
		if recover() != nil {
			send = true
		}
	}()
	c <- value
	return false
}
func main() {
	c := make(chan TT, 2)
	////c <- 1 //如果有数据发送,则判断不正确.
	//fmt.Println(IsClose(c))
	//close(c)
	//fmt.Println(IsClose(c))
	c <- 1
	//fmt.Println(SaveIsClose(c, 1))
	close(c)
	//fmt.Println(SaveIsClose(c, 1))
	debug.PrintStack()
}

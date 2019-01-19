package main

import (
	"os"
	"os/signal"
	"fmt"
)
func main() {
	StudySignal()
}
/*******************
* signal 是监听信号的,如果收不到信息则一直阻塞
 */
func StudySignal() {
	ch := make(chan os.Signal)
	//signal.Notify(ch, os.Interrupt, os.Kill)
	signal.Notify(ch)
	s := <- ch
	fmt.Println(s)
}

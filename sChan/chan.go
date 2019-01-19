package sChan

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

/**********************
*
 ***********************/
func StudyChan() {
	stop := make(chan bool)
	go func() {
		for {
			fmt.Println(<-stop)
			select {
			case <-stop:
				fmt.Println("关闭系统")
			default:
				time.Sleep(time.Second)
				fmt.Println("正在运行...")
			}
		}
	}()
	log.Log().Str("开始监控", "Ok")
	time.Sleep(time.Second * time.Duration(3))
	stop <- true
	time.Sleep(time.Second * time.Duration(10))
}

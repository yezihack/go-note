/*
1. 为什么需要协程池?
控制协程池的数量, 复用每一个协程,执行更多任务.

1. 创建一个协程池, 控制协程的数量
2. 协程池里创建一个队列,不断取数据,让正在运行的协程去执行他们
3. 外部也需要一个队列,负责接受任务数据
4. 外部的队列与内部的队列需要建立一个通道,有任务就转移给内部队列

*/
package main

import (
	"fmt"
	"time"
)

//创建一个task任务
type Task struct {
	f func() error
}

//将任务执行掉
func (t *Task) Execute() {
	t.f()
}

//创建一个任务
func NewTask(task func() error) *Task {
	return &Task{
		f: task,
	}
}

//定义一个池子
type Pool struct {
	JobsChan chan *Task //内部队列
	WaitChan chan *Task //外部队列
	WorkNum  int        //控制协程数量
}

//创建一个池子
func NewPool(cap int) *Pool {
	return &Pool{
		JobsChan: make(chan *Task),
		WaitChan: make(chan *Task),
		WorkNum:  cap,
	}
}

//让池子里的任务工作起来
func (p *Pool) work(workID int) {
	for task := range p.JobsChan {
		task.Execute()
		fmt.Println("workID", workID)
	}
}

//让协程池跑起来
func (p *Pool) run() {
	for i := 0; i < p.WorkNum; i++ {
		go p.work(i)
	}
	//将外部列队里的任务转移到内部队列来
	for task := range p.WaitChan {
		p.JobsChan <- task
	}
}

func main() {
	//创建任务
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	//实例一个池子.
	pool := NewPool(3)
	//加入外部队列
	i := 0
	go func() {
		for {
			pool.WaitChan <- t
			i++
			fmt.Println("第", i, "个任务")
		}
	}()

	//运行协程池
	pool.run()

}

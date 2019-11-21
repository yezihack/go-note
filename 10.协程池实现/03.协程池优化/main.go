package main

import (
	"fmt"
	"runtime"
	"time"
)

//定义一个任务结构,装载任务的
type Task struct {
	task func()
}

//消费任务
func (t *Task) Execute() {
	t.task()
}

//实例一个任务对象.
func NewTask(InTask func()) *Task {
	t := Task{
		task: InTask,
	}
	return &t
}

//定义一个协程池结构
type Pool struct {
	jobQueue   chan *Task    //内部队列
	EntryQueue chan *Task    //外部队列
	cap        int           //协程数量
	Sign       chan struct{} //设置一个信号灯
	totalWork  int           //处理了多少任务
}

//实例一个pool对象.
func NewPool(cap int) *Pool {
	p := Pool{
		jobQueue:   make(chan *Task),
		EntryQueue: make(chan *Task),
		cap:        cap,
		Sign:       make(chan struct{}),
	}
	p.run()
	return &p
}

//让协程池里的任务工作起来
func (p *Pool) work(workID int) {
	for job := range p.jobQueue {
		job.Execute()
		p.totalWork++
	}
}

//让协程池工作起来, 并创建一个调度goroutine
func (p *Pool) run() {
	for i := 0; i < p.cap; i++ {
		go p.work(i)
	}
	//将外部队列转移给内部队列
	go p.dispatch()
	go p.close()
}

//调度, 将外部的任务调度到内部的任务队列里去
func (p *Pool) dispatch() {
	for t := range p.EntryQueue {
		p.jobQueue <- t
	}
}

//等待信号量传入, 如果收到则关闭队列的通道
func (p *Pool) close() {
	fmt.Println("正在关闭")
	<-p.Sign
	close(p.EntryQueue)
	close(p.jobQueue)
	fmt.Println("已经关闭")
	p.Monitor()
}

//通知关闭
func (p *Pool) Release() {
	fmt.Println("发出信号")
	p.Sign <- struct{}{}
}

//添加监控goroutine数量
func (p *Pool) Monitor() {
	fmt.Printf("共执行了%d个任务\n", p.totalWork)
	fmt.Printf("当前的goroutine数量:%d\n", runtime.NumGoroutine())
}

func main() {
	p := NewPool(5)
	for i := 0; i < 100; i++ {
		p.EntryQueue <- NewTask(func() {
			//	fmt.Println(time.Now())
		})
	}
	p.Release()
	time.Sleep(time.Second * 1)
}

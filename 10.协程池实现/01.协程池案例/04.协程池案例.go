package main

import "fmt"

//1.创建一个任务结构,用于装载任务使用.
//2.创建一个Pool任务池,含一个chan队列
//3.开辟指定大小的goroutine
type Task struct {
	taskFunc func()
}

func (t *Task) execute() {
	t.taskFunc()
}
func NewTask(inTask func()) *Task {
	return &Task{
		taskFunc: inTask,
	}
}

type Pool struct {
	TaskQueue chan *Task //task queue list
	cap       int
}

func NewPool(QueueNumber int) *Pool {
	p := Pool{
		TaskQueue: make(chan *Task),
		cap:       QueueNumber,
	}
	p.run()
	return &p
}
func (p *Pool) work(workID int) {
	for task := range p.TaskQueue {
		task.execute()
		fmt.Printf("第:%d个工作在干活\n", workID)
	}
}
func (p *Pool) run() {
	for i := 1; i <= p.cap; i++ {
		go p.work(i)
	}
}
func (p *Pool) Release() {
	close(p.TaskQueue)
}

func main() {
	p := NewPool(5)
	defer p.Release()
	for i := 1; i <= 100; i++ {
		i := i
		p.TaskQueue <- NewTask(func() {
			fmt.Println("任务:", i)
		})
	}

}

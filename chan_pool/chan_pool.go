package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

//实现一个简单的协和池

//定义一个池子
type Pool struct {
	JobQueue          chan Job  //定义一个任务队列
	WorkCurrentNumber int       //当前工作的协程数
	MaxWorker         int       //设置最大的协程数
	Result            chan bool //标记是否已经完成工作
}

//定义一个任务结构
type Job struct {
	ID int
}

//定义一个工作结构
type Worker struct {
	Result chan bool
}

//开始工作
func (w *Worker) DoJob(job Job) {
	fmt.Printf("正在进行%d号工作\n", job.ID)
	time.Sleep(time.Millisecond * 500)
	w.Result <- true
}

//向工作队列里添加工作
func (p *Pool) AddJob(job Job) {
	p.JobQueue <- job
}

//监听工作是否已经完成
func (p *Pool) stop() {
	for {
		<-p.Result
		p.WorkCurrentNumber--
	}
}
func (p *Pool) Run() {
	go p.stop()
	for {
		if p.WorkCurrentNumber < p.MaxWorker { //如果队列还没满则继续处理工作
			select {
			case job, ok := <-p.JobQueue: //取出一个工作任务进行处理
				if ok {
					worker := &Worker{
						Result: p.Result,
					}
					go worker.DoJob(job) //处理工作
					p.WorkCurrentNumber++
				} else {
					log.Print("jobQuery is over")
					return
				}
			}
		}
	}
}
func main() {
	jobQueue := make(chan Job)
	result := make(chan bool)
	p := &Pool{
		JobQueue:  jobQueue,
		Result:    result,
		MaxWorker: 20,
	}
	go func() {
		for i := 0; i < 100; i++ {
			job := Job{i}
			p.AddJob(job) //添加任务
		}
		close(jobQueue)
	}()
	p.Run()
}

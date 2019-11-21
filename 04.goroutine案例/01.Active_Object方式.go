package main

import (
	"fmt"
	"sort"
	"sync"
)
//思路: 开多个goroutine向chan管道里放数据.保证开一个goroutine去写数据
//这样就避免高并发带来的资源竞争问题.
//重点是开一个goroutine写数据,没有资源争抢.
//开多个goroutine向管道里添加数据,管道是并发安全的,保证高效性.
type Service struct {
	channel chan int `desc:"即将加入到数据slice的数据"`
	data []int `desc:"数据slice"`
}
//新建一个size大小的缓冲Active Object对象
func NewService(num int, done func()) *Service {
	s := &Service{
		channel:make(chan int, num),
		data :make([]int, 0),
	}
	go func() {
		s.schedule()
		done()
	}()
	return s
}
//将管道里的数据添加到slice里
func (s *Service) schedule() {
	for v := range s.channel  {
		s.data = append(s.data, v)
	}
}
//向chan里添加值.
func (s *Service) Add(i int) {
	s.channel <- i
}
//用完关闭管道
func (s *Service) Close() {
	close(s.channel)
}
//获取slice
func (s *Service) Slice() []int {
	sort.Ints(s.data)
	return s.data
}

func main() {
	//新建一个信号
	c := make(chan struct{})
	s := NewService(2, func() {
		c <- struct{}{}
	})
	//开启n个goroutine添加数据
	n := 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0;i < n; i ++ {
		go func(a int) {
			s.Add(a)
			wg.Done()
		}(i)
	}
	wg.Wait()
	s.Close()
	<-c
	//检查数据是否完整

	fmt.Printf("个数:%d, 结果:%v\n", len(s.Slice()), s.Slice())
}

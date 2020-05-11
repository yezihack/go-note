package sChan

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"math"
	"runtime"
	"sync"
	"unsafe"
)

/**********************
*
 ***********************/
func StudyChan() {
	stop := make(chan bool)
	go func() {
		for {
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
	time.Sleep(time.Second * time.Duration(1))
	stop <- true
	time.Sleep(time.Second * time.Duration(2))
}

func GoWait() {
	var wg sync.WaitGroup
	wg.Add(1)
	//go func() {
	//	wg.Wait()
	//	fmt.Println("exit")
	//}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("done")
		wg.Done()
	}()
	wg.Wait()
	wg.Wait()
	fmt.Println("exit2")
}
func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}
	fmt.Println(x)
}
func CreateCount(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}
func GoCount(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}
	wg.Wait()
}

func GoSched() {
	c := make(chan struct{})
	go func() {
		defer close(c)
		go func() {
			println("first go")
		}()
		for i := 0; i < 5; i++ {
			if i == 3 {
				runtime.Gosched() //暂停,等待下一次调度
			}
			println("i:", i)
			runtime.Goexit()
		}
		go func() {
			fmt.Println("abc")
		}()
	}()
	<-c
	fmt.Println("exit 0")
}

//通道只是一个队列, 同步模式下, 发送和接受方必须配对.否则阻塞.
//异步模式, 缓冲未满时不会阻塞
func SignChan() {
	c := make(chan byte)
	exit := make(chan struct{})
	go func() {
		c <- 'c'
		close(exit)
	}()
	fmt.Println(<-c)
	<-exit
}

func CacheChan() {
	c := make(chan string, 3)
	exit := make(chan struct{})

	c <- "a"
	c <- "a"
	c <- "a"
	close(exit)

	fmt.Println(<-c)
	fmt.Println(len(c), cap(c), unsafe.Sizeof(c))
	i := math.MaxUint32
	str := "abc"
	fmt.Println(i, unsafe.Sizeof(i), unsafe.Sizeof(str))
	<-exit
}

func ForChan() {
	c := make(chan int)
	done := make(chan struct{})
	go func() {
		//使用 for来消费chan
		for x := range c {
			println(x)
		}
		close(done)
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-done
}

func RandChan() {
	var wg sync.WaitGroup
	wg.Add(2)
	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()
		for {
			var (
				val int
				ok  bool
			)
			select {
			case val, ok = <-a:
				fmt.Println(val)
			case val, ok = <-b:
				fmt.Println(val)
			}
			if !ok { //如果任意一个通道关闭,则终止接收
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)
		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()
	wg.Wait()

}
func RandDefaultChan() {
	done := make(chan struct{})
	data := []chan int{
		make(chan int, 3),
	}
	go func() {
		defer close(done)
		for i := 0; i < 10; i++ {
			select {
			case data[len(data)-1] <- i: //向通道放入数据.
			default: //缓冲满了则创建一个
				data = append(data, make(chan int, 3))
			}
		}
	}()
	<-done
	//读取数据.
	for i := 0; i < len(data); i++ {
		tmp := data[i]
		close(tmp)
		for x := range tmp {
			fmt.Println(x)
		}
	}
}

// 工厂方法与goroutine绑定
type receive struct {
	sync.WaitGroup
	data chan int
}

func NewReceive(n int) *receive {
	r := &receive{
		data: make(chan int),
	}
	r.Add(1)
	go func(n int) {
		r.Done()
		for x := range r.data {
			fmt.Printf("i: %d, val:%d \n", n, x)
		}
	}(n)
	return r
}

func ShowReceive(n int) {
	r := NewReceive(n)
	r.data <- 10
	r.data <- 20
	//close(r.data) //关闭通道 , 发出结束通知
	r.Wait() //等待接收者处理结束
}

// goroutine 实现一个安全队列
type pool chan []byte

func NewPool(cap int) pool {
	return make(chan []byte, cap)
}
func (p pool) Put(b []byte) {
	select {
	case p <- b:
	default:
	}
}
func (p pool) Get() []byte {
	var v []byte
	select {
	case v = <-p:
	default:
		v = make([]byte, 10)
	}
	return v
}
func ShowNewPool() {
	p := NewPool(10)
	p.Put([]byte("aaa"))
	v := p.Get()
	fmt.Println(string(v))
}

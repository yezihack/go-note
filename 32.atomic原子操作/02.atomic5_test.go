package _2_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAddInt32(t *testing.T) {
	var a, b int32
	for i := 0; i < 3; i ++ {
		atomic.AddInt32(&a, 3) //3相当于步长.必须获取a的地址才能进行递增.
		atomic.AddInt32(&b, -3) //3相当于步长.必须获取b的地址才能进行递增.
	}
	fmt.Println(a, b) //9, -9
}
func TestCAS(t *testing.T) {
	var o, n int32
	o = 3
	n = 10
	//atomic.CompareAndSwapInt64()
	//atomic.CompareAndSwapPointer()
	//atomic.CompareAndSwapUint32()
	//atomic.CompareAndSwapUint64()
	//atomic.CompareAndSwapUintptr()
	if atomic.CompareAndSwapInt32(&o, 3, n) {
		fmt.Println("cas交换成功", o, n)
	} else {
		fmt.Println("cas交换失败", o, n)
	}
}
var value int32
func AddValue(delta int32) {
	for {
		v := value
		//值进行比较.
		if atomic.CompareAndSwapInt32(&value, v, v + delta) {
			fmt.Println("exit CAS")
			break//退出自旋锁.
		}
	}
}

func AddValueV2(delta int32) {
	for {
		//v := value //操作会有并发安全问题. 需要对称修改
		//原子的读取value的值, 把它赋值给变量v
		//这样一来, 当前计算机中任何cpu都不会进行其他针对此值的读写操作,
		//这样的约束到底层硬件的支持.
		v := atomic.LoadInt32(&value)
		// 值进行比较.
		if atomic.CompareAndSwapInt32(&value, v, v + delta) {
			fmt.Println("exit CAS")
			break//退出自旋锁.
		}
	}
}
func TestCASFOR(t *testing.T) {
	value = 5
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		delay := time.NewTimer(time.Second * 2)
		select {
		case <-delay.C:
			value = 5
			fmt.Println("value:", value)
			return
		}
	}()
	go func() {
		defer wg.Done()
		AddValue(10)
	}()
	wg.Wait()
	fmt.Println("value", value)
}
func Timer() {
	d := time.Millisecond * 500
	delay := time.NewTimer(d)
		delay.Reset(d)
		select {
		case <-delay.C:
			fmt.Println("时间到了")
		}
}
func TestTimer(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		Timer()
	}()
	wg.Wait()

}
func TestStore(t *testing.T) {
	var v int32
	atomic.StoreInt32(&v, 10)
	fmt.Println(atomic.LoadInt32(&v))
}
func TestSwap(t *testing.T) {
	var v int32
	v = 30
	//直接设置赋值操作,一定会成功,返回旧值.
	old := atomic.SwapInt32(&v, 10)
	fmt.Println(old, v) //30, 10
}
func TestValue(t *testing.T) {
	//如果第一次存储的类型, 第二次再次存储时,也要是同样的类型, 否则会出现panic
	//atomic.Value不能做为参数传入函数.虽然编辑不会出错, 这样做会导致并发不安全.因为会产生副本.
	var v atomic.Value
	v.Store(func() {
		fmt.Println("atomic value")
	})
	v.Store(func() { //第二次存储时,必须也第一次存储的类型一致, 否则panic
		fmt.Println("hello world")
	})
	fn := v.Load().(func())
	fn()
	fn = v.Load().(func())
	fn()
}
func FuncValue(v atomic.Value) {
	i := v.Load().(int)
	fmt.Println(i)
}

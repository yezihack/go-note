package sync_study

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//sync.Cond 条件变量, 条件变量是构建在一个基础锁上的同步原语
/*
L 锁操作
Wait 挂起goroutine (必须获取该锁之后才能调用Wait()方法)
Signal 唤醒一个goroutine
Broadcast 唤醒所有的goroutine

唤醒因wait condition而挂起goroutine，区别是Signal只唤醒一个，
而Broadcast唤醒所有。允许调用者获取基础锁Locker之后再调用唤醒，但非必需。

必须获取该锁之后才能调用Wait()方法，Wait方法在调用时会释放底层锁Locker，并且将当前goroutine挂起，
直到另一个goroutine执行Signal或者Broadcase，该goroutine才有机会重新唤醒，并尝试获取Locker，完成后续逻辑。

1）Cond 不能被复制。
因为 Cond 内部维护着一个所有等待在这个 Cond 的 Go 程队列，如果 Cond 允许值传递，则这个队列在值传递的过程中会进行复制，导致在唤醒 Go 程的时候出现错误。

（2）唤醒顺序。
从等待队列中按照顺序唤醒，先进入等待队列，先被唤醒。
*/
var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func TestCond(t *testing.T) {
	for i := 0; i < 10; i ++ {
		go func(x int) {
			cond.L.Lock() //获取锁
			defer cond.L.Unlock()//释放锁
			cond.Wait() ////等待通知，阻塞当前 goroutine
			//todo
			fmt.Println("x", x)
		}(i)
	}
	go func() {
		fmt.Println("我行我素")
	}()


	time.Sleep(time.Second) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态


	fmt.Println("1Signal...")
	cond.Signal()               // 1 秒后下发一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("2Signal...")
	cond.Signal()               // 1 秒后下发下一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("3Signal...")
	cond.Signal()               // 1 秒后下发下一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("Broadcast...")
	cond.Broadcast()            // 1 秒后下发广播给所有等待的goroutine

	time.Sleep(time.Second * 1)	// 睡眠 1 秒，等待所有 goroutine 执行完毕
}
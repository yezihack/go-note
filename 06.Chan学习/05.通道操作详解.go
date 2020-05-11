package main

func ChanClose() {
	var ch chan int
	close(ch) //关闭一个零值nil的通道, 会引起panic
	//解决以上问题.见ChanCloseV2
}
func ChanCloseV2() {
	var ch chan int
	ch = make(chan int)
	close(ch) //关闭一个非零值nil的通道,属于正常关闭
}
func ChanCloseClose() {
	ch := make(chan int) //非零值nil的通道
	close(ch)
	close(ch) //关闭一个已关闭的通道会panic
	//解决方法见ChanCloseCloseV2
}
func ChanCloseCloseV2() {
	ch := make(chan int) //非零值nil的通道
	close(ch) //对一个通道只能关闭一次.
}
func ChanUnInit() {
	var ch chan int
	ch <- 1 //向一个零值nil通道发送数据, 永久阻塞
}
func ChanNotNilClose() {
	ch := make(chan int)
	close(ch)
	ch <- 1 //向已关闭的通道发送数据会panic
}
func ChanNotNilUnClose() {
	ch := make(chan int)
	ch <- 1 //向一个未关闭的通道发送数据, 会阻塞或成功, 主要取消有没有另一端接收数据.
}
func ChanNil() {
	var ch chan int
	<- ch
}
func ChanReceiveInitClosed() {
	ch := make(chan int)
	close(ch)
	<- ch //一个已关闭的通道, 接收数据, 永不阻塞
}
//对于剩下的情况（即此通道的缓冲队列和发送数据协程队列均为空），
// 此接收数据协程Gr将被推入接收数据协程队列，并进入阻塞状态。
// 它以后可能会被另一个发送数据协程唤醒而恢复运行。
// 对于这种情况，此数据接收操作为一个阻塞操作。
func ChanReceiveInitUnClose() {
	ch := make(chan int)
	<- ch //一个未关闭的通道， 接收数据， 阻塞 或者成功接收
}
// 如果此通道的缓冲队列不为空（这种情况下，接收数据协程队列必为空），
// 此协程Gr将从缓冲队列取出接收一个值。 如果发送数据协程队列不为空，
// 一个发送协程将从此队列中弹出，此协程欲发送的值将被推入缓冲队列。
// 此发送协程将恢复至运行状态。 接收数据协程Gr继续运行，不会阻塞。
// 对于这种情况，此数据接收操作为一个非阻塞操作。
func ChanReceiveInitUnCloseV2() {
	ch := make(chan int, 1) //缓冲通道
	ch <- 1
	<- ch //一个未关闭的通道， 接收数据， 阻塞 或者成功接收
}
//否则（即此通道的缓冲队列为空），如果发送数据协程队列不为空（这种情况下，此通道必为一个非缓冲通道），
// 一个发送数据协程将从此队列中弹出，此协程欲发送的值将被接收数据协程Gr接收。
// 此发送协程将恢复至运行状态。 接收数据协程Gr继续运行，不会阻塞。
// 对于这种情况，此数据接收操作为一个非阻塞操作。
func ChanReceiveInitUnCloseV3() {
	ch := make(chan int) //缓冲通道
	go func() { //发送协程队列不为空
		ch <- 1
	}()
	<- ch //一个未关闭的通道， 接收数据， 阻塞 或者成功接收
}
func main() {
	//ChanClose() //panic
	//ChanCloseV2() //ok
	//ChanCloseClose()//panic
	//ChanCloseCloseV2()//ok
	//ChanUnInit() //永久阻塞, deadlock
	//ChanNotNilClose() //panic
	//ChanNotNilUnClose() //永久阻塞, deadlock
	//ChanNil()//永久阻塞, deadlock
	//ChanReceiveInitClosed() //ok
	//ChanReceiveInitUnClose()//永久阻塞, deadlock
	//ChanReceiveInitUnCloseV2()//ok
	//ChanReceiveInitUnCloseV3() //ok
	//ChanNotNilClose()
}
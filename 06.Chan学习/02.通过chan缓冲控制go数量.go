package main

type Number struct {
	Ch           chan int //通道
	Data         []int    //装载数据
	CurrentNum   int      //当前数量
	MaxGoroutine int      //设置起多少个goroutine
}

func NewNumber(max int) *Number {
	return &Number{
		Ch:           make(chan int, max),
		Data:         make([]int, 0),
		MaxGoroutine: max,
	}
}

func (n *Number) PushServe() {
	for c := range n.Ch {
		go func() {
			n.Push(c)
		}()
	}
}
func (n *Number) Push(i int) {
	n.Data = append(n.Data, i)
}
func main() {
	NewNumber(10)

}

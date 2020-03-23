
# CONTEXT 上下文
> 控制goroutine的生命周期.

## Background 与 TODO 什么时候使用
```go
// 一般main, init使用
func Background() Context {
	return background
}
// 新建context接口或不清楚时就使用TODO()
func TODO() Context {
	return todo
}
```

## Context衍生四个函数
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```
> func WithCancel(parent Context) (ctx Context, cancel CancelFunc) 

WithCancel函数是最常用的，传入父context，它会返回一个子context接口和
一个用来通知取消子context的函数；

> func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

WithDeadline函数除了需要传入父context，还需要传入截止时间，也是返回一个子context接口和一个用来通知取消子context的函数，到了指定的截止时间会自动取消，我们也可以使用取消函数提前取消；

> func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

WithTimeout函数和WithDeadline函数类似，只是传入的是超时的时间，指多少时间后自动取消，而不是指定的时间点，即一个是时间长度，一个是时间点。

> func WithValue(parent Context, key, val interface{}) Context

WithValue函数不是和取消相关的，而是通过context在goroutine之间传递一个key-value对，可以通过上面说的context.Value()方法获取。
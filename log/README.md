# 错误与异常

### 错误panic
- panic 会导致程序中止运行
```
func main() {
	//直接导致程序中止运行
	panic("出错啦")
}
```
- 使用recover可以防止程序中止运行并友好输出错误
```
func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("不会导致程序中止运行")
		}
	}()
	//直接导致程序中止运行
	panic("出错啦")
}
```
- 如果是协程,recover必须在协程里才会起作用

```
func main() {
	MyGoroutine()
	time.Sleep(time.Second)
}
func MyGoroutine() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("不会导致程序中止运行")
		}
	}()
	//直接导致程序中止运行
	panic("出错啦")
}

```
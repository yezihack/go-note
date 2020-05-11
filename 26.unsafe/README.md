# unsafe 指针

## go指针的限制
1. 限制一：Go 的指针不能进行数学运算
```
a := 5
p := &a
p++
p = &a + 3
```
2. 限制二：不同类型的指针不能相互转换。
```
func main() {
    a := int(100)
    var f *float64

    f = &a
}
```
3. 限制三：不同类型的指针不能使用 == 或 != 比较
> 只有在两个指针类型相同或者可以相互转换的情况下，才可以对两者进行比较。另外，指针可以通过 == 和 != 直接和 nil 作比较
4. 限制四：不同类型的指针变量不能相互赋值。

## 为什么有 unsafe
1. Go 语言类型系统是为了安全和效率设计的，有时，安全会导致效率低下。
有了 unsafe 包，高阶的程序员就可以利用它绕过类型系统的低效。
因此，它就有了存在的意义，阅读 Go 源码，会发现有大量使用 unsafe 包的例子

## unsafe 实现原理
unsafe 包提供了 2 点重要的能力：
1. 任何类型的指针和 unsafe.Pointer 可以相互转换。
1. uintptr 类型和 unsafe.Pointer 可以相互转换。

uintptr 
1. uintptr 并没有指针的语义，意思就是 uintptr 所指向的对象会被 gc 无情地回收
1. unsafe.Pointer 有指针语义，可以保护它所指向的对象在“有用”的时候不会被垃圾回收

关于unsafe.Pointer的4个规则。

任何指针都可以转换为unsafe.Pointer
1. unsafe.Pointer可以转换为任何指针
1. uintptr可以转换为unsafe.Pointer
1. unsafe.Pointer可以转换为uintptr
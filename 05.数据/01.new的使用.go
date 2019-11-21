/*
new 是分配内存的内建函数
它不会初使化内存,只会将内存置零.
也就是说:New(T)会为类型T的新项分配已置零的空间,并返回它的地址,也就是一个类型为*T的值
简称返回一个指针.该指针指向新分配的类型为T的零值.
*/
package main

import "fmt"

func main() {
	i := new(int)
	*i = 10
	fmt.Printf("addr:%p, value:%d\n", i, *i)
	//再次使用new, 结果应该是重新分配一个零值的内存空间
	i = new(int)
	fmt.Printf("addr:%p, value:%d\n", i, *i)
	//addr:0xc000058080, value:10
	//addr:0xc000058098, value:0 地址被重新分配.与上面的地址不一致
}

package main

import "fmt"
//如果切片的容量小于1024个元素，那么扩容的时候slice的cap就翻番，乘以2；
// 一旦元素个数超过1024个元素，增长因子就变成1.25，即每次增加原来容量的四分之一。
//如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组，
// 如果扩容之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
func main() {
	//初例一个0长度的切片,容量默认
	s := make([]int, 1) //按照长度的2倍扩容
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))//len:3, cap:4
	oldCap := cap(s)
	//如果元素超过1024个. 扩容则以1.25,即1/4容量扩容.
	for i := 0;i < 2000; i ++ {
		s = append(s, i)
		if cap(s) != oldCap {
			oldCap = cap(s)
			fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))//len:3, cap:4
		}
	}

}

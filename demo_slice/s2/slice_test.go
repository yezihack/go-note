package s2

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	list := [3]int{1, 2, 3} //定义3个数
	fmt.Printf("len:%d, cap:%d, list:%v\n", len(list), cap(list), list)
	list1 := list[0:1] //浅表复制 给list1 浅拷贝同根
	list1[0] = 100     //修改list1第一个元素, list的第一个元素也会被修改
	fmt.Printf("len:%d, cap:%d, list:%v\n", len(list), cap(list), list)
	fmt.Printf("len1:%d, cap1:%d, list1:%v\n", len(list1), cap(list1), list1)
	list2 := list[0:2]                        //浅表复制给list2
	list2 = append(list2, []int{200, 300}...) //扩容摆脱同根, 切片追加, 将开辟一个新的空间,将元素复制过
	list2[1] = 800                            //向切片追加大于cap大小, 将之前的元素复制给list2开辟的空间.并修改,将不会修改list1
	fmt.Printf("len:%d, cap:%d, list:%v\n", len(list), cap(list), list)
	fmt.Printf("len1:%d, cap1:%d, list1:%v\n", len(list1), cap(list1), list1)
	fmt.Printf("len2:%d, cap2:%d, list2:%v\n", len(list2), cap(list2), list2)
}

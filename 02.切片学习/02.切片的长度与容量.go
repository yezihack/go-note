package main

import "fmt"

//切片的长度与容量

func main() {
	//创建一个切片, 长度为2, 容量为4
	//长度为2,相当于先创建2个空间.
	//容量为4,相当于我一共有4个空间.
	//如果使用的长度超出容量长度,按2倍扩容.
	s1 := make([]int, 2, 4)
	//向s1添加1个数据.查看长度与容量的变化
	s1 = append(s1, 5)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))//len:3, cap:4

	//再向s1添加2个数据, 也就是说一共有5个数据, 大于总容量的大小. 看长度与容量的变化
	s1 = append(s1, 6, 7)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))//len:5, cap:8, 结果就是按2倍大小扩容.
}

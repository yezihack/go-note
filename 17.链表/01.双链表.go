package _7_链表

import (
	"container/list"
	"fmt"
)
//官方自带container/list包, 提供双链表功能
func Linked() {
	l := list.New()//初始一个链表
	for i := 0; i < 5; i ++ {
		l.PushBack(i)//向链表中添加元素, 向后添加
	}
	head := l.Front() //获取头指针
	PrintLinked(head)
	fmt.Println("打印尾指针数据", l.Back().Value) //打印尾指针数据
}

func PrintLinked(head *list.Element) {
	str := ""
	for head != nil {
		str += fmt.Sprint(head.Value)//打印当前结点数据
		head = head.Next()//下一个结点
		if head != nil {
			str += fmt.Sprint("->")
		}
	}
	fmt.Println(str)
}
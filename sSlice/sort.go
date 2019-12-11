package sSlice

import (
	"fmt"
	"sort"
)

type Person struct {
	Age int
}
type Man struct {
	Age int
}
//定义切片类型
type Mans []Man

//实现排序len函数
func (m Mans) Len() int {
	return len(m)
}
//进行升序 i大于j 为降序, i小于j 为升序
func (m Mans) Less(x, y int) bool {
	return m[x].Age > m[y].Age
}
//swap 进行位置置换
func (m Mans) Swap(x, y int) {
	m[x], m[y] = m[y], m[x]
}
func LessSort() {
	//声明切片对象
	person := Mans{
	}
	//添加值
	person = append(person, Man{
		Age: 10,
	})
	person = append(person,Man{
		Age: 2,
	})
	person = append(person, Man{
		Age: 15,
	})

	fmt.Println("未排序:")
	for _, item := range person{
		fmt.Printf("Age: %d\n", item.Age)
	}
	sort.Sort(person)
	fmt.Println("IS Sorted:", sort.IsSorted(person))
	fmt.Println("排序后-降序:")
	for _, item := range person {
		fmt.Printf("Age: %d\n", item.Age)
	}
}

func SliceSort() {
	//声明切片对象
	person := make([]*Person, 0)
	//添加值
	person = append(person, &Person{
		Age: 10,
	})
	person = append(person, &Person{
		Age: 2,
	})
	person = append(person, &Person{
		Age: 15,
	})
	person = append(person, &Person{
		Age: 29,
	})
	person = append(person, &Person{
		Age: 9,
	})
	fmt.Println("未排序:")
	for _, item := range person{
		fmt.Printf("Age: %d\n", item.Age)
	}
	//进行降序 i大于j 为降序
	sort.Slice(person, func(i, j int) bool {
		return person[i].Age > person[j].Age
	})
	fmt.Println("排序后-降序:")
	for _, item := range person{
		fmt.Printf("Age: %d\n", item.Age)
	}

	//进行升序 i < j
	sort.Slice(person, func(i, j int) bool {
		return person[i].Age < person[j].Age
	})
	fmt.Println("排序后-升序:")
	for _, item := range person{
		fmt.Printf("Age: %d\n", item.Age)
	}
}

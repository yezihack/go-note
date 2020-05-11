/*
映射map,是一种key,value存储结构.也是一种引用类型
使用前必须make分配内存空间
获取一个不存在的map键值,会返回该映射中项的类型对应的零值.
*/
package main

import "fmt"

func main() {
	//1. 声明一个map
	m := make(map[int]string)
	//2. 赋值操作
	m[10] = "wang"
	//3.获取操作
	fmt.Println(m[10])
	//4. 获取一个不存在的key
	fmt.Println(m[11])//获取一个不存在的map键值,会返回该映射中项的类型对应的零值.
	//5. 判断一个不存在的key
	_, ok := m[1]
	if ok {
		fmt.Println("key存在")
	} else {
		fmt.Println("key不存在")
	}
	//6. 删除key
	delete(m, 10)
	fmt.Println(m)
	//7. 打印
	m[0] = "king"
	fmt.Printf("%+v\n", m)
	fmt.Printf("%#v\n", m)
}

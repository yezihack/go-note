package _0_基础

import (
	"fmt"
	"testing"
)
type User struct {
	Name string
}
func TestMap(t *testing.T) {
	my := make(map[int]User, 100)
	my[0] = User{"golang"}
	my[1] = User{"php"}
	//从 map 中取回的是⼀一个 value 临时复制品，对其成员的修改是没有任何意义的。
	//my[0].Name = "c"//cannot assign to struct field my[0].Name in map
	//正确做法是完整替换 value 或使⽤用指针。
	u := User{"c++"}
	my[0] = u
	for _, v := range my {
		fmt.Println(v.Name)
	}
	my2 := map[int]*User{
		1:&User{"python"},
	}
	my2[1].Name = "c#"
	fmt.Println(my2)
}

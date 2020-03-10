package method

import (
	"fmt"
	"testing"
)

//1. 对自定义的变量进行添加方法
type Money float64
func (m *Money) Format() string {
	return fmt.Sprintf("%.2f", *m)
}
//2. 对结构体添加变量
type Student struct {
	Name string
}
//
func (s Student) Print() {
	fmt.Println("name:", s.Name)
}

func (s *Student) SetName(name string){
	s.Name = name
}

func TestMain1(t *testing.T) {
	var f Money = 1.24342
	fmt.Println(f.Format())

	 s := Student{"baili"}
	 s.Print()
}

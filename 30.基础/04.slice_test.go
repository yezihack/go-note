package _0_基础

import (
	"fmt"
	"testing"
)

//需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引⽤用数组⽚片段，以
//实现变⻓长⽅方案。

//可直接创建 slice 对象，⾃自动分配底层数组
func TestSlice(t *testing.T) {
	a := []int{2,3, 4}
	fmt.Printf("len: %d, cap:%d, %v\n", len(a), cap(a), a)
}

func TestReSlice(t *testing.T) {
	//所谓 reslice [x:y:z] x是包括元素位置, y是不包括元素位置, z是虚拟位置,即cap, x:z之间
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a1 := a[1:3]
	a2 := a[2:5:6]
	a2[0] = 100 //slice是引用类型, 会修改a, a1, a2切片都会被修改
	fmt.Printf("a len:%d, cap:%d, %v\n", len(a), cap(a), a)//len: 9, cap: 9, [1 2 3 4 5 6 7 8 9]
	fmt.Printf("a1 len:%d, cap:%d, %v\n", len(a1), cap(a1), a1)//len:2, cap: 8, [2,3]
	fmt.Printf("a2 len:%d, cap:%d, %v\n", len(a2), cap(a2), a2)//len:3, cap: 4, [3, 4, 5]
}
//⼀一旦超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
func TestAppend(t *testing.T) {
	a := []int{1,2,3}
	a = append(a, 4, 5)
}
//函数 copy 在两个 slice 间复制数据，复制⻓长度以 len ⼩小的为准。两个 slice 可指向同⼀一
//底层数组，允许元素区间重叠
func TestCopy(t *testing.T) {
	arr := [...]int{1,2,3,4,5,6,7,8}
	a1 := arr[:5]//[1 2 3 4 5]
	a2 := arr[5:]//[6 7 8]
	fmt.Println(a1, a2)
	copy(a2, a1)//[1 2 3 4 5] [1 2 3]
	fmt.Println(a1, a2)
	a2 = []int{6, 7, 8}
	copy(a1, a2) //
	fmt.Println(a1, a2)
	//copy(dst, src),将src里的元素复制给dst里去.
	c := make([]int, len(a1))
	copy(c, a1)
	c[0] = 100
	fmt.Println("c", a1, c)

}
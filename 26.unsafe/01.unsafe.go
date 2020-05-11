package _6_unsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)
/*

*/

//unsafe直接绕过类型检查,操作底层指针
func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bb := reflect.SliceHeader{
		Data:stringHeader.Data,
		Len:stringHeader.Len,
		Cap:stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bb))
}
func byte2string(b []byte) string{
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.SliceHeader{
		Data:sliceHeader.Data,
		Len:sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func Double(x *int) {
	*x += *x
	x = nil
	//但是对 x 本身（一个指针）的操作却不会影响外层的 a，所以 x = nil 掀不起任何大风大浪。
}


func Int2float(i int) {
	//获取int的地址
	i = 10
	iptr := &i
	//通过unsafe.Pointer转换成Float指针类型
	var fp *float64 = (*float64)(unsafe.Pointer(iptr))
	//对fp指针值的操作,i变量也会同样被改变
	*fp = *fp * 2
	fmt.Println(i)

}
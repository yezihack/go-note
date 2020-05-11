package sConst

import "fmt"

const (
	Unknow = iota
	Man
	Gril
	AA = iota
	BB
	CC = iota + 2
	DD = iota
	EE
)

func GetVal() {
	fmt.Println("Man", Man)
	fmt.Println("DD", DD)
	fmt.Println("EE", EE)
}

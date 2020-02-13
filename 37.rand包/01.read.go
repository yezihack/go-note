package main

import (
	"fmt"
	"math/rand"
)

//”crypto/rand“ 包实现了用于加解密的更安全的随机数生成器。
func main() {
	for i := 0 ; i< 10;i ++ {
		bb := make([]byte, 10)
		n, _ := rand.Read(bb)
		fmt.Println(n, bb, string(bb))
		for _, b := range bb {
			fmt.Print(string(b))
		}
		fmt.Println()
	}
}

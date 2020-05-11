package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

//16进制解码
func HexDecode(s string) []byte {
	dst := make([]byte, hex.DecodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.DecodedLen
	n, err := hex.Decode(dst,  []byte(s))//进制转换, src->dst
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return dst[:n] //返回0:n的数据.
}
//字符串转为16进制
func HexEncode(s string) []byte {
	dst := make([]byte, hex.EncodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.EncodedLen
	n := hex.Encode(dst, []byte(s)) //字节流转化成16进制
	return dst[:n]
}

func main() {
	s16 := "6769746875622e636f6d2f79657a696861636b"
	fmt.Println(string(HexDecode(s16)))

	s := "github.com/yezihack"
	fmt.Println(string(HexEncode(s)))

	fmt.Println(hex.Dump([]byte(s)))
}

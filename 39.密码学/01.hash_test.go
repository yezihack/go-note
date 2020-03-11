package _9_密码学

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

//使用io写入字符串，计算出sha1 sum值
func TestHash(t *testing.T) {
	hasher := sha1.New() //声明一个sha1对象
	io.WriteString(hasher, "test") //将字符串写入hasher对象里
	b := make([]byte, 0)
	result1 := hasher.Sum(b)
	fmt.Println("sha1 sum:", fmt.Sprintf("%x", result1))
	//a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
	fmt.Println("sha1 sum:", fmt.Sprintf("%d", result1))
	//[169 74 143 229 204 177 155 166 28 76 8 115 211 145 233 135 152 47 187 211]
}
func TestSha1Write(t *testing.T) {
	h := sha1.New()
	h.Write([]byte("lesmoon"))
	result := h.Sum(nil)
	fmt.Println("sha1 sum:", fmt.Sprintf("%x", result))
	//1dbdd20f852e0730c704f8373a33dd647c52f8d4
}

package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

// 完全读取 response.body 数据
// 必须将 http.Response 的 Body 读取完毕并且关闭后，才会重用底层的 TCP 连接
func ReadBodyAll(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	// 完全读取 response.body
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
}
// 不读取 response.body 数据
func NotReadBody(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	// 不读取 response.body  只有关闭没有读取 body 的话, tcp不会重用.
	// _, err = ioutil.ReadAll(resp.Body)
}
func main() {
	count := 100
	for i := 0; i <; i++ {

	}
}
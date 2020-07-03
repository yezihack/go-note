/**
 * @Author WANGZILIANG
 * @date 2020/6/30 11:44
 * @Project_name yezihack
 */
package main

import "testing"
var (
	TestURL = "https://www.cnblogs.com/52php/p/6985411.html"
)

func BenchmarkReadBodyAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadBodyAll(TestURL)
	}
}

func BenchmarkNotReadBody(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotReadBody(TestURL)
	}
}
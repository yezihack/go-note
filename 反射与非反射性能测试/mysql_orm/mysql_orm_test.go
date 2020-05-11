package mysql_orm

import (
	"testing"
)

func TestGetBlackListByOri(t *testing.T) {
	//fmt.Println(GetBlackListByOri())
	//fmt.Println(GetBlackListByOrm())
	//InsertBlackList()
}

//func BenchmarkGetListByOri(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		GetListByOri()
//	}
//}
//func BenchmarkGetListByOrm(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		GetListByOrm()
//	}
//}

func BenchmarkGetBlackListByOri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetBlackListByOri()
	}
}
func BenchmarkGetBlackListByOrm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetBlackListByOrm()
	}
}

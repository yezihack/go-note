package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)


func TestPointer(t *testing.T) {
	type mystr struct {
		name *string
	}
	s := new(mystr)
	name := "china"
	s.name = &name
	fmt.Println(*s.name)
}

func BenchmarkAdd(b *testing.B) {
	Add(1000)
}
func BenchmarkFb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fb(20)
	}
}
func TestFb2(t *testing.T) {
	rs := Fb2(365)
	fmt.Println(rs)
}
func BenchmarkFb2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fb2(1000)
	}
}
func TestTime(t *testing.T) {
	s := "18:30"
	fmt.Println(SubMinute(s))
}

//计算时间差, 正则相离, 负则已过
func SubMinute(HourMinute string) interface{} {
	var BeijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)
	Ymd := "2006-01-02"
	YmdAll := "2006-01-02 15:04:05"
	now := time.Now()
	start := now.Format(Ymd) + " " + HourMinute + ":00"
	s, err := time.ParseInLocation(YmdAll, start, BeijingLocation)
	if err != nil {
		return 0
	}
	return int(s.Sub(now.In(BeijingLocation)).Minutes())
}
func TestFile(t *testing.T) {
	a := CheckDirExists("/Users/wangzl/go-work/src/github.com/yezihack/gofun/")
	fmt.Println(a)
}

//判断文件是否存在
func CheckFileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

//判断文件夹是否存在
func CheckDirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func TestNow(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	t1 := time.Now().Add(1).UnixNano()
	fmt.Println(t1)

}

func TestLine(t *testing.T) {
	fmt.Println(len("big_data_donation_share_friend_shareto微信分享关系好友,携带是否是朋友圈好友"))

}

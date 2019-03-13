package main

import (
	"fmt"

	"bytes"
	"regexp"
	"strings"
	"sync"

	"log"
	"reflect"
)

func main() {
	//buff := new(bytes.Buffer)
	//log.SetOutput(buff)
	//log.SetFlags(log.Lshortfile)
	//log.SetPrefix("[INFO]")
	//log.Print("test")
	//log.SetOutput(os.Stdout)
	//log.Print("aaaa")
	//buff.Reset()
	var r []int
	p(r)
}
func p(value interface{}) {
	dest := reflect.Indirect(reflect.ValueOf(value))
	fmt.Println(dest.Kind())
	fmt.Println(reflect.ValueOf(value).Kind())
}

func tt1() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile|log.LstdFlags)
		infof = func(info string) {
			logger.Output(2, info)
		}
	)

	infof("Hello world")

	fmt.Print(&buf)
}
func testlog() {
	buff := new(bytes.Buffer)
	log.New(buff, "ERROR", log.LstdFlags|log.Lshortfile)
	log.Print("aaaaabbbbbbbbbbb")
}

type one struct {
}

var oc sync.Once
var oe *one

func NewOne() *one {
	oc.Do(func() {
		oe = &one{}
	})
	return oe
}

func CheckCharDoSpecialArr(s string, char byte) []string {
	s = CheckCharDoSpecial(s, char)
	return strings.Split(s, string(char))
}

//检查字符串,去掉特殊字符, 仅支持大小写字母,特殊字符: _ - , 三种,其它一律过滤掉
func CheckCharDoSpecial(s string, char byte) string {
	reg := regexp.MustCompile(`[a-zA-Z\_\-\,0-9]`)
	var result []string
	if arr := reg.FindAllString(s, -1); len(arr) > 0 {
		fmt.Println(arr)
		buf := bytes.Buffer{}
		for key, val := range arr {
			if val != string(char) {
				buf.WriteString(val)
			}
			if val == string(char) && buf.Len() > 0 {
				result = append(result, buf.String())
				buf.Reset()
			}
			//处理最后一批数据
			if buf.Len() > 0 && key == len(arr)-1 {
				result = append(result, buf.String())
			}
		}
	}
	return strings.Join(result, string(char))
}

//拼接特殊字符串
func FormatField(field string, formats []string) string {
	buf := bytes.Buffer{}
	for key := range formats {
		buf.WriteString(fmt.Sprintf(`%s:"%s" `, formats[key], field))
	}
	return strings.TrimRight(buf.String(), " ")
}
func InArrayString(str string, arr []string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}

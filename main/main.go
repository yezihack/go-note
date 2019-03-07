package main

import (
	"fmt"

	"bytes"
	"strings"
	"regexp"
	"sync"
)

func main() {
	s1 := NewOne()
	s2 := NewOne()
	if s1 == s2 {
		fmt.Println("相同")
	} else {
		fmt.Println("不相同")
	}
}

type one struct {
}
var oc sync.Once
var oe *one
func NewOne() *one{
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
			if buf.Len() > 0 && key == len(arr) - 1 {
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
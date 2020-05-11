package main

import (
	"fmt"

	"bytes"
	"regexp"
	"strings"
	"sync"

	"log"
	"reflect"
	"compress/gzip"
	"io"
	"time"
	"os"
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
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	// Setting the Header fields is optional.
	zw.Name = "a-new-hope.txt"
	zw.Comment = "an epic space opera by George Lucas"
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte("A long time ago in a galaxy far, far away..."))
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

}

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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

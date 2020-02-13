package file_operation

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

//文件操作.

//将文件整个读入内存
func GetFileAllContent(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func TestGetFileAllContent(t *testing.T) {
	filename := "./data.txt"
	data, err := GetFileAllContent(filename)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
//第二种方法,读取整个文件.ioutil.ReadFile封装了.os.open方法
func GetFileAllContentV2(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func TestGetFileAllContentV2(t *testing.T) {
	filename := "./data.txt"
	data, err := GetFileAllContentV2(filename)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

package file_operation

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

//只读整个文件的部分内容.
func ReadByteFileContent(filename string) (data string, err error){
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	//只读取10个字节.
	buf := make([]byte, 3)
	count, err := file.Read(buf)
	if err != nil {
		return
	}
	fmt.Println("all", string(buf))
	fmt.Println("part", string(buf[:count]))
	return
}
func TestReadByteFileContent(t *testing.T) {
	filename := "./data.txt"
	data, err := ReadByteFileContent(filename)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func ReadLineFileName(filename string) (data string, err error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(stat)
	var size = stat.Size()
	fmt.Println("size", size)
	//把字节流读取bufio
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println("line", line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read is end")
			}
			break
		}
	}
	return
}
func TestReadLineFileName(t *testing.T) {
	filename := "./data.txt"
	data, err := ReadLineFileName(filename)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
//http://www.jquerycn.cn/a_21681
//设置文件偏移量, 读取文件内容
func ReadLineFileNameV2(filename string) (data string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	stat, _ := file.Stat()
	fmt.Println(stat.Size())
	count, err := file.Seek(0, os.SEEK_CUR)
	if err != nil {
		return
	}
	fmt.Println("count", count)
	buf := make([]byte, 10)
	_, err = file.Read(buf)
	fmt.Println(string(buf))
	return "", nil
}
func TestReadLineFileNameV2(t *testing.T) {
	filename := "./data.txt"
	data, err := ReadLineFileNameV2(filename)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
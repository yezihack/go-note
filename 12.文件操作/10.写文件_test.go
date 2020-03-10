package file_operation

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"
	"unsafe"
)

func WriteFileNumber(data, path string) (bool, error) {
	file, err := os.Create(path)
	if err != nil {
		return false, err
	}
	count, err := file.WriteString(data)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, errors.New("count" + strconv.Itoa(count))
}
func TestWriteFileNumber(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	for i := 0; i < 10; i ++ {
		buf.WriteString(strconv.Itoa(i))
	}
	_, err := WriteFileNumber(buf.String(), "./data.txt")
	if err != nil {
		t.Error(err)
	}
}
func TestSizeOf(t *testing.T) {
	var i int32
	i = 10
	fmt.Println(unsafe.Sizeof(i))
	var s struct{}
	fmt.Println(unsafe.Sizeof(s))
	st := struct {
		name string
		//age int
		//index int
	}{}
	fmt.Println(unsafe.Sizeof(st))
}

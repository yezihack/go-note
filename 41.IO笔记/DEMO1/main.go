package main

import (
	"bytes"
	"github.com/unknwon/com"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		body := request.Body
		data := CopyData(body)
		writer.Write(data)
	})
}

func demo1() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)
	io.Copy(os.Stdout, lr)
}

func CopyData(rd io.Reader) []byte  {
	lr := io.LimitReader(rd, 5)
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, lr)
	return buf.Bytes()
}
func CopyTee(rd io.Reader) io.Writer {
	buf := bytes.NewReader(nil)
	io.TeeReader(buf, )
}

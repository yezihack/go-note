package main

import "fmt"

type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}
type Streamer interface {
	Write(buf []byte) (n int, err error)
	Read(buf []byte) (n int, err error)
}
type File struct {
}
func (f *File) Write(buf []byte) (n int, err error) {
	return 0, nil
}
func (f *File) Read(buf []byte) (n int, err error) {
	return 0, nil
}
func (f *File) Show() {
}
type Writer interface {
	Write(buf []byte) (n int, err error)
}
func main() {
	var f = File{}
	var r ReadWriter = &f
	var i Streamer = &f
	var w Writer = &f
	r.Read([]byte("hello"))
	i.Read([]byte("hello"))
	w.Write([]byte("ok"))
	if f2, ok := w.(Streamer); ok {
		fmt.Println(f2.Write([]byte{'a'}))
		fmt.Println(f2.Read([]byte{'a'}))
	}
}

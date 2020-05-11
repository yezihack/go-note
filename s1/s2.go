package main

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}
func (s *S) Write(str string) {
	s.data = str
}

func main() {
	m := map[int]S{1: {"A"}}
	m[1].Read()
	m[1].Write("B")
}

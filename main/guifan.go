package main

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}
func (s *S) Write(data string) {
	s.data = data
}
func main() {
	sVals := map[int]S{1: {"A"}}
	sVals[1].Read()
	sVals[1].Write("B")
}

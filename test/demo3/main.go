package main

import (
	"fmt"
	"time"
)

type Car struct {
	names []string
}

func (d *Car) SetNames(n []string) {
	d.names = n
}
func (d *Car) SetNames2(n []string) {
	nn := make([]string, len(n))
	copy(nn, n)
	d.names = nn
}

func (d *Car) Show() {
	fmt.Println(d.names)
}

func At(n time.Duration) {
	//todo
	time.Sleep(n)
}

func main() {
	s := []string{"a", "b"}
	c := Car{}
	c.SetNames2(s)
	c.Show()
	s[0] = "ccc"
	c.Show()

}

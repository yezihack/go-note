package main

import (
	"code.qschou.com/golang/errors"
	"fmt"
	"log"
	"math"
)

type PathErr struct {
	Name string
	Err error
}
func (e *PathErr) Error() string {
	return "Err:" + e.Name + e.Err.Error()
}

func sqrt(i int) (result int, err error) {
	if i <= 0 {
		return 0, &PathErr{"sqrt", errors.New("i < 0")}
	}
	return int(math.Sqrt(float64(i))), nil
}

func main() {
	ret, err := sqrt(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}

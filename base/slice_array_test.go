package _base_test

import (
	"fmt"
	"testing"
	"time"
)

func TestSliceArray(t *testing.T) {
	a := make([]int, 5)
	b := make([]int, 5, 8)
	for i := 0; i < 5; i++ {
		if len(a) > i {
			a = append(a, i)
			continue
		}
		a[i] = i
	}
	println(len(a), cap(a))
	fmt.Println(a)
	println(len(b), cap(b))
}
func TestSlice(t *testing.T) {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Println(s3[0:2]) //相当于[0,2) 下标, 2不包括
}
func TestChan(t *testing.T) {
	c := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go func() {
			for x := range c {
				fmt.Println(x)
			}
		}()
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func TestChan2(t *testing.T) {
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func(j int) {
			c <- j
		}(i)
	}
	for x := range c {
		fmt.Println(x)
	}
	close(c)
}

func TestChan3(t *testing.T) {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(j int) {
			c <- j
		}(i)
	}

	go Display(c)
	time.Sleep(time.Millisecond * 500)
}

func Display(c chan int) {
	for x := range c {
		fmt.Println(x)
	}
}

func TestChan4(t *testing.T) {
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func(j int) {
			c <- j
		}(i)

	}
STOP:
	for {
		select {
		case x, ok := <-c:
			if ok {
				fmt.Println(x)
			}
		case <-time.After(time.Millisecond * 500):
			break STOP
		}
	}
}

func TestArr(t *testing.T) {
	a := [3]int{1, 2, 3}
	a[2] = 10
	fmt.Println(a)

	arr := [5]int{}
	arr[0] = 100
	fmt.Println(arr)

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 [2]int
	arr2[0] = 1
	fmt.Println(arr2)

	var arr3 = [...]int{}
	fmt.Println(arr3)

}

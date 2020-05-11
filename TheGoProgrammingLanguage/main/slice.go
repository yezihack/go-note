package main

import "fmt"

func main() {
	myMap := make(map[string]int, 10)
	fmt.Println(myMap)
	delete(myMap, nil)
}
func smap() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6, 7, 8}
	s3 := make([]int, 2)
	//copy(s1, s2)//s1: [4 5 6]
	//copy(s2, s1)//s2: [1 2 3 7 8]
	copy(s3, s1)// s3:[1,2]
	fmt.Println(s1, s2, s3)

	str1 := []string{"aa", "bb"}
	str2 := []string{"cc", "dd"}
	copy(str1, str2)//str1: [cc dd] [cc dd]
	fmt.Println(str1, str2)
}
func scap() {
	s := make([]int, 0, 4)
	for i := 0; i < 2048; i ++ {
		s = append(s, i)
		if cap(s) % 512 == 0 {
			fmt.Printf("i:%d, len:%d, cap:%d\n",i, len(s), cap(s))
		}
	}
}

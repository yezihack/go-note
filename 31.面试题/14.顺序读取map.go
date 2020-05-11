package main

import (
	"fmt"
	"sort"
)

func main() {
	name := []string{"a", "wo", "kk", "pp", "bb"}
	sort.Strings(name)
	fmt.Println(name)
}

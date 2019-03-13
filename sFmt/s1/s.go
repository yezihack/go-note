package s1

import "fmt"

func GetS1() {
	type aa struct {
		name string
	}
	b := make([]aa, 0)
	b = append(b, aa{
		name: "king",
	})
	b = append(b, aa{
		name: "wang",
	})
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%#v\n", b)
}

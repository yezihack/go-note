package public

import "fmt"

type Public struct {

}
func (p *Public) show() {
	fmt.Println("public-show")
}

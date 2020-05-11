package main

import "fmt"

type Carer interface {
	Show()
	Input(s string)
}

type Car struct {
}

func (*Car) Show() {
	fmt.Println("Car interface")
}
func (Car) Input(s string) {
	fmt.Println(s)
}

type Car1 struct {
}

func (*Car1) Show() {
	fmt.Println("Car interface")
}
func (*Car1) Input(s string) {
	fmt.Println(s)
}

type Car2 struct {
}

func (Car2) Show() {
	fmt.Println("Car interface")
}
func (Car2) Input(s string) {
	fmt.Println(s)
}

func GetCar(c Carer) Carer {
	return c
}

func main() {
	var c Car
	var cc Carer = &c //*Car才复合Carer要求
	cc.Show()
	cc.Input("Car0")

	var c1 Car1
	var cc1 Carer = &c1 //*Car才复合Carer要求
	cc1.Show()
	cc1.Input("Car1")

	var c2 Car2
	var cc2 Carer = c2 //*Car才复合Carer要求
	cc2.Show()
	cc2.Input("Car2")
	fmt.Println("-------------------------")
	GetCar(&c).Show()
	GetCar(&c).Input("tCar")

}

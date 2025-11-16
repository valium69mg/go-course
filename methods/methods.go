package main

import "fmt"

type Rectangle struct {
	base   float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.base * r.height
}

func (r *Rectangle) Scale(factor float64) {
	r.base *= factor
	r.height *= factor
}

type MyInt int

func main() {
	rectangle := Rectangle{base: 20, height: 20}
	fmt.Println(rectangle.Area())
	rectangle.Scale(2)
	fmt.Println(rectangle)
	fmt.Println(rectangle.Area())

	myInt := MyInt(2)
	fmt.Println(myInt.isPositive())
	fmt.Println(myInt.welcome())
}

func (m MyInt) isPositive() bool {
	return m > 0
}

func (MyInt) welcome() string {
	return "Welcome!"
}

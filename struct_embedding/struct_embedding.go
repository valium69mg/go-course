package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person     // embedded struct
	employeeId int
	salary     float64
}

func (p person) introduce() {
	fmt.Println("Hello, my name is:", p.name)
}

// overrides the inherited introduce from person
func (e employee) introduce() {
	fmt.Printf("Hello, my name is: %s, my employee id is: %d\n", e.name, e.employeeId)

}

func main() {

	employee := employee{person: person{name: "Bob", age: 35}, employeeId: 1, salary: 50.000}
	fmt.Println(employee)
	fmt.Println(employee.name)
	employee.introduce()
}

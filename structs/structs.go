package main

import "fmt"

type Address struct {
	street string
	number int
}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
}

func main() {

	person := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	fmt.Println(person)

	person2 := Person{
		firstName: "John",
	}

	fmt.Println(person2)

	fmt.Println(person.firstName)

	personAddress := &person

	fmt.Printf("Person address %p\n", personAddress)

	fmt.Println(*personAddress)

	personAddress.firstName = "Johnny"

	fmt.Println(*personAddress)
	fmt.Println(person)

	fmt.Println(person.fullName())

	user := struct {
		username string
		password string
	}{
		username: "johndoe123",
		password: "lol122",
	}

	fmt.Println(user)

	person.incrementAge(1)

	fmt.Println(person)

	person.address = Address{
		street: "Heaven street",
		number: 33,
	}

	fmt.Println(person)

}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAge(years int) {
	p.age += years
}

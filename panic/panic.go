package main

import "fmt"

func main() {
	// panic
	process(10)

	process(-10)

}
func process(input int) {

	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")

	if input >= 0 {
		defer fmt.Println("Before panic")
		fmt.Printf("processing input: %d\n", input)
	}
	panic("Input must be a non negativa number")
}

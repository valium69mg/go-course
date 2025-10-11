package main

import (
	"errors"
	"fmt"
)

func main() {
	// functions that starts with upper case are public, and private start with lowercase
	var num int = 2
	fmt.Println(doubleInt(num))

	fmt.Println(add(num, num))

	// anonymous funcions or closure funcions
	greet := func() {
		fmt.Println("Hello anonymous!")
	}

	greet()

	func() {
		fmt.Println("Hello again, anonymous!")
	}()

	// functions as types
	operations := add

	resultOperation := operations(num, num)

	fmt.Println(resultOperation)

	// func as arg
	fmt.Println(applyOperation(num, num, operations))

	// function that returns a function
	multiplers := createMultiplier(5)

	fmt.Println(multiplers(num))

	// multiple return values
	a := 17
	b := 3
	fmt.Println(divide(a, b))

	// errors
	for i := 1; i < 5; i++ {
		result, error := compare(1, i)
		if error != nil {
			fmt.Printf("Error: %s\n", error)
		} else {
			fmt.Println(result)
		}
	}

}

func doubleInt(num int) int {
	return num * 2
}

func add(a, b int) int {
	return a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is grater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	}
	return "", errors.New("Error: numbers are equal")
}

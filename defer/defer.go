package main

import "fmt"

func main() {

	// defer statements invokes a function whose execution is deferred to the moment the sorrounding function returns

	process(1)

}

func process(i int) {
	defer fmt.Printf("Deffered i value: %d\n", i)
	defer fmt.Println("Deffered statement excecuted first!")
	defer fmt.Println("Deffered statement excecuted second!")
	i++
	fmt.Println("Normal excecution statement")
	defer fmt.Printf("i value: %d\n", i)

}

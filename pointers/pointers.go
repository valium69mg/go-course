package main

import "fmt"

func main() {
	var ptr *int
	var a int = 10
	ptr = &a
	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	modifyValue(ptr)
	fmt.Println(*ptr)

}

func modifyValue(ptr *int) {
	*ptr++
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \n Golang!"
	rawMessage := `Hello, \n Golang!` // this is raw, escape sequences do not work!
	fmt.Println(message)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable:", len(message))

	fmt.Println("The first character of message is:", message[0]) // this will print in ASCII

	greeting := "hello, "
	name := "carlos"
	fmt.Println(greeting + name)

	apple := "Apple"
	var banana string = "Banana"
	fmt.Println(apple > banana)

	for index, char := range apple {
		fmt.Printf("Index: %d, char: %c, hex: %x, ASCII: %v\n", index, char, char, char)
	}

	// runes
	fmt.Println("Rune count in string: ", utf8.RuneCountInString(apple))

	var ch rune = 'C'
	fmt.Printf("Rune of character %c: %d\n", ch, ch)

	var cstring string = string(ch)
	fmt.Println(cstring)

	fmt.Printf("Type of cstring is: %T, type of ch is: %T\n", cstring, ch)
}

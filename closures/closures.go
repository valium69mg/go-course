package main

import "fmt"

func main() {

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())

	substractor := func() func(int) int {
		countdown := 90
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(substractor(1))
	fmt.Println(substractor(1))
	fmt.Println(substractor(1))
	fmt.Println(substractor(1))

}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i: ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

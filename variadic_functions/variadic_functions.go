package main

import "fmt"

func main() {
	// ... Ellipsis
	// func functionName(param1 type1, param2 type2, param3 ...type3) returnType {
	//	body
	//}

	// variadic parameters can accept 0 or more parameters of a certain type

	fmt.Println(sum(1, 2, 3, 4, 5))

	message, total := sumMultiReturn("Total is:", 1, 2, 3, 4, 5)
	fmt.Printf("%s %d\n", message, total)

	// variadic parameter must be the last parameter
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	message2, total2 := sumMultiReturn("Total is:", numbers...) // spread operator

	fmt.Printf("%s %d\n", message2, total2)

}

func sum(nums ...int) int {
	var total int = 0
	for _, v := range nums {
		total += v
	}
	return total
}

func sumMultiReturn(returnString string, nums ...int) (string, int) {

	total := 0
	for _, v := range nums {
		total += v
	}

	return returnString, total
}

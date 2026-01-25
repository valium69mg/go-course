package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &myError{code: -1, message: "Square root method not allowed for negative values."}
	}
	return math.Sqrt(x), nil
}

type myError struct {
	code    int
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func main() {
	x := 2.0
	result, err := sqrt(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	y := -1.0
	result, err = sqrt(y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

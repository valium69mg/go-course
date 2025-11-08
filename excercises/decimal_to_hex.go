package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	number := 531814
	numberToHex := decimalToHex(number)
	fmt.Println(numberToHex)
}

func decimalToHex(number int) string {
	hex_chars := []string{}
	for {
		remainder := number % 16
		if remainder == 0 {
			break
		}
		remainderToHex, err := numberToHexCharacter(remainder)
		if err != nil {
			panic(err)
		}
		hex_chars = append(hex_chars, remainderToHex)
		number /= 16
	}
	hex_chars_str := strings.Join(hex_chars, "")
	return reverseString(hex_chars_str)
}
func numberToHexCharacter(number int) (string, error) {
	if number < 0 || number > 15 {
		return "", fmt.Errorf("Hex number range must be between 0 and 15, out of range value: %d", number)
	}

	if number >= 0 && number < 10 {
		return strconv.Itoa(number), nil
	} else {
		switch number {
		case 10:
			return "A", nil
		case 11:
			return "B", nil
		case 12:
			return "C", nil
		case 13:
			return "D", nil
		case 14:
			return "E", nil
		case 15:
			return "F", nil
		}
	}
	return "", fmt.Errorf("Unknown error ocurred, no convertion was found for number: %d", number)

}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

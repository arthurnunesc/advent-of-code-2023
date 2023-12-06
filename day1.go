package main

import (
	"fmt"
	"strconv"
	"strings"
)

// function to reverse string
func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func getFirstDigit(str string) (result rune) {
	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			return ch
		}
	}
	return
}

func getLastDigit(str string) (result rune) {
	for _, ch := range reverseString(str) {
		if ch >= '0' && ch <= '9' {
			return ch
		}
	}
	return
}

func getFirstAndLastNumberAsInt(str string) (result int) {
	first_digit := getFirstDigit(str)
	last_digit := getLastDigit(str)
	// fmt.Printf("first_digit = %c\n", first_digit)
	// fmt.Printf("last_digit = %c\n", last_digit)
	current_number, _ := strconv.Atoi(string(first_digit) + string(last_digit))
	return current_number
}

func main() {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
kasjdkasjdklsaj`
	var total int = 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		total += getFirstAndLastNumberAsInt(line)
	}
	// fmt.Println("no of lines:", len(lines))
	fmt.Println("total =", total)
}

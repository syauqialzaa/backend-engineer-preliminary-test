package algorithm

import (
	"fmt"
	"strings"
)

// Function to calculate 10 raised to the power of the number of digits minus 1
func PowerOf10(num int) int {
	power := 1
	for num >= 10 {
		power *= 10
		num /= 10
	}
	return power
}

func PseudoCode() {
	inputNumber := "1.345.679"

	// Remove dots and convert the string to an integer
	inputNumber = strings.ReplaceAll(inputNumber, ".", "")
	num := 0
	fmt.Sscanf(inputNumber, "%d", &num)

	// Process each digit
	for num > 0 {
		// Get the last digit
		digit := num % 10

		// Output the digit in the desired format
		fmt.Println(digit * PowerOf10(num))

		// Remove the last digit
		num /= 10
	}
}

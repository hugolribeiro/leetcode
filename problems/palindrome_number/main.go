package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome(545))
	fmt.Println(isPalindrome(9807))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversedX := reverseANumber(x)
	return reversedX == x
}

func reverseANumber(number int) int {
	reversedNumber := 0
	amountDigits := getAmountDigitsFromANumber(number)
	for i := amountDigits; i > 0; i-- {
		reversedNumber += (number % 10) * int((math.Pow10(i - 1)))
		number /= 10
	}
	return reversedNumber
}

func getAmountDigitsFromANumber(number int) int {
	digits := 1
	for {
		number /= 10
		if number >= 1 {
			digits += 1
		} else {
			break
		}
	}
	return digits
}

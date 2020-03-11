package main

import (
	"fmt"
)

func main() {
	getNthFib(10)
}

func getNthFib(n int) int {
	var firstNumber = 0
	var secondNumber = 1
	var nextNumber int
	var counter = 3

	if n == 1 {
		return firstNumber
	}

	if n == 2 {
		return secondNumber
	}

	for counter <= n {
		nextNumber = firstNumber + secondNumber
		firstNumber = secondNumber
		secondNumber = nextNumber

		counter++
	}

	fmt.Println(n, "th Fib:", secondNumber)
	return secondNumber
}

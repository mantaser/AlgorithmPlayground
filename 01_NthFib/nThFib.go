package main

import (
	"fmt"
)

// The Fibonacci sequence is defined as follows: the first number of the sequence is 0, the second number is 1,
// and the nth number is the sum of the (n - 1)th and (n - 2)th numbers. Write a function that takes in an integer n
// and returns the nth Fibonacci number.

// Important note: the Fibonacci sequence is often defined with its first 2 numbers as F0 = 0 and F1 = 1.
// For the purpose of this question, the first Fibonacci number is F0; therefore, getNthFib(1) is equal to F0, getNthFib(2) is equal to F1, etc..

// Sample input #2: 6
// Sample output #2: 5 (0, 1, 1, 2, 3, 5)

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

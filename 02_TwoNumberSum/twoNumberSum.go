package main

import (
	"fmt"
)

// Write a function that takes in a non-empty array of distinct integers and an integer representing a target sum.
// If any two numbers in the input array sum up to the target sum, the function should return them in an array.
// If no two numbers sum up to the target sum, the function should return an empty array.
// Assume that there will be at most one pair of numbers summing up to the target sum.

// Sample input: [3, 5, -4, 8, 11, 1, -1, 6], 10
// Sample output: [-1, 11]

func main() {
	inputArray := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10

	fmt.Println("First solution:", twoNumberSum(inputArray, targetSum))
	fmt.Println("Second solution:", twoNumberSumHashTable(inputArray, targetSum))
}

// Time O(n^2) Space O(1)
func twoNumberSum(array []int, target int) []int {
	arrayLength := len(array)
	index := 0
	var sum int

	for index < arrayLength-1 {
		innerIndex := index + 1

		for innerIndex < arrayLength {
			sum = array[index] + array[innerIndex]

			if sum == target {
				return []int{array[index], array[innerIndex]}
			}

			innerIndex++
		}

		index++
	}

	return []int{}
}

// Time O(n) Space O(n)
func twoNumberSumHashTable(array []int, target int) []int {
	valuesHashtable := map[int]bool{}

	for _, value := range array {
		pairNumber := target - value
		_, exists := valuesHashtable[pairNumber]

		if exists {
			return []int{pairNumber, value}
		}

		valuesHashtable[value] = true
	}

	return []int{}
}

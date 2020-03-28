package main

import (
	"math"
)

// BST is a Binary Search Tree struct
type BST struct {
	value int

	left  *BST
	right *BST
}

// FindClosestValue will find the closest given target value in the BST structure
func (tree *BST) FindClosestValue(target int) int {
	currNode := tree
	closestValue := getClosestValue(currNode.value, math.MaxInt32, target)

	for currNode != nil {
		if currNode.value < target {
			currNode = currNode.right
		} else {
			currNode = currNode.left
		}

		if currNode != nil {
			closestValue = getClosestValue(closestValue, currNode.value, target)
		}
	}

	return closestValue
}

func getClosestValue(value int, comparedValue int, target int) int {
	if absInt(value-target) <= absInt(comparedValue-target) {
		return value
	}

	return comparedValue
}

func absInt(value int) int {
	if value > 0 {
		return value
	}

	return -value
}

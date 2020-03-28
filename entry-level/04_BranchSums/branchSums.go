package main

// BinaryTree is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func main() {

}

// BranchSums is a fuction to count tree branches sums
func BranchSums(root *BinaryTree) []int {
	var currSum int
	branchSums := make([]int, 0)
	return getBranchSums(root, branchSums, currSum)
}

func getBranchSums(node *BinaryTree, branchSums []int, currSum int) []int {
	currSum += node.Value

	if node.Left != nil {
		branchSums = getBranchSums(node.Left, branchSums, currSum)
	}

	if node.Right != nil {
		branchSums = getBranchSums(node.Right, branchSums, currSum)
	}

	if node.Left == nil && node.Right == nil {
		branchSums = append(branchSums, currSum)
	}

	return branchSums
}

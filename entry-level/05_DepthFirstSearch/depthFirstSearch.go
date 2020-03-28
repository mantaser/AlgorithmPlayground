package main

// Node is a tree like node struct
type Node struct {
	Name     string
	Children []*Node
}

func main() {

}

// DepthFirstSearch returns a string array of depth first search nodes names
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)

	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}

	return array
}

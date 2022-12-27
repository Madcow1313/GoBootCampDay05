package main

import (
	"fmt"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func walkThroughTree(bt *TreeNode) int {
	sum := 0
	if bt == nil {
		return 0
	}
	if bt.HasToy {
		sum += 1
	}
	if bt.Left != nil {
		sum += walkThroughTree(bt.Left)
	}
	if bt.Right != nil {
		sum += walkThroughTree(bt.Right)
	}
	return sum
}

func areToysBalanced(bt *TreeNode) bool {
	if bt == nil || (bt.Left == nil && bt.Right == nil) {
		return true
	}
	sumLeft := walkThroughTree(bt.Left)
	sumRight := walkThroughTree(bt.Right)
	return sumLeft == sumRight
}

func createNode(hasToy bool) *TreeNode {
	tree := new(TreeNode)
	tree.HasToy = hasToy
	tree.Left = nil
	tree.Right = nil
	return tree
}

func main() {
	/*
		     0
			/ \
		   0   1
		  / \
		 0   1
	*/
	firstRoot := createNode(false)
	firstRoot.Left = createNode(false)
	firstRoot.Left.Left = createNode(false)
	firstRoot.Left.Right = createNode(true)
	firstRoot.Right = createNode(true)

	fmt.Printf("Expected for first tree <%t>, reality <%t>\n", true, areToysBalanced(firstRoot))

	/*
			   1
		     /  \
		    1     0
		   / \   / \
		  1   0 1   1

	*/

	secondRoot := createNode(true)
	secondRoot.Left = createNode(true)
	secondRoot.Left.Left = createNode(true)
	secondRoot.Left.Right = createNode(false)
	secondRoot.Right = createNode(false)
	secondRoot.Right.Left = createNode(true)
	secondRoot.Right.Right = createNode(true)

	fmt.Printf("Expected for second tree <%t>, reality <%t>\n", true, areToysBalanced(secondRoot))

	/*
			  1
		     / \
		    1   0

	*/
	thirdRoot := createNode(true)
	thirdRoot.Left = createNode(true)
	thirdRoot.Right = createNode(false)
	fmt.Printf("Expected for third tree <%t>, reality <%t>\n", false, areToysBalanced(thirdRoot))

	/*
			  0
		     / \
		    1   0
		     \   \
		      1   1

	*/

	fourthRoot := createNode(false)
	fourthRoot.Left = createNode(true)
	fourthRoot.Left.Right = createNode(true)
	fourthRoot.Right = createNode(false)
	fourthRoot.Right.Right = createNode(false)
	fmt.Printf("Expected for fourth tree <%t>, reality <%t>\n", false, areToysBalanced(fourthRoot))

	/*
			1
		   /
		  1
	*/

	firstTestRoot := createNode(true)
	firstTestRoot.Left = createNode(true)
	fmt.Printf("Expected for first test tree <%t>, reality <%t>\n", false, areToysBalanced(firstTestRoot))

	/*
		1
	*/

	secondTestRoot := createNode(true)
	fmt.Printf("Expected for second test tree <%t>, reality <%t>\n", true, areToysBalanced(secondTestRoot))

	fmt.Println("Crash test", areToysBalanced(nil))

}

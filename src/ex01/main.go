package main

import (
	"fmt"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

// func leftAndRightAppend(root *TreeNode) []bool {
// 	answer := make([]bool, 0)
// 	if root.Left != nil {
// 		answer = append(answer, root.Left.HasToy)
// 	}
// 	if root.Right != nil {
// 		answer = append(answer, root.Right.HasToy)
// 	}
// 	return answer
// }

func reverseBoolSlice(slice []bool) []bool {
	var output []bool
	for i := len(slice) - 1; i >= 0; i-- {
		output = append(output, slice[i])
	}
	return output
}

// func recForwardWalk(root *TreeNode) []bool {
// 	answer := make([]bool, 0)
// 	if root == nil {
// 		return answer
// 	}
// 	answer = append(answer, leftAndRightAppend(root)...)
// 	answer = append(answer, recForwardWalk(root.Left)...)
// 	answer = append(answer, recForwardWalk(root.Right)...)
// 	return answer
// }

func getCurrentLevel(root *TreeNode, level int) []bool {
	answer := make([]bool, 0)
	if root == nil {
		return answer
	}
	if level == 1 {
		answer = append(answer, root.HasToy)
		return answer
	} else {
		answer = append(answer, getCurrentLevel(root.Left, level-1)...)
		answer = append(answer, getCurrentLevel(root.Right, level-1)...)
	}
	return answer
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func unrollGarland(root *TreeNode) []bool {

	answer := make([]bool, 0)
	height := getHeight(root)
	for i := 0; i <= height; i++ {
		if i%2 != 0 {
			answer = append(answer, reverseBoolSlice(getCurrentLevel(root, i))...)
		} else {
			answer = append(answer, getCurrentLevel(root, i)...)
		}
	}
	// answer = append(answer, root.HasToy)
	// answer = append(answer, recForwardWalk(root)...)
	return answer
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
			1
		   /  \
		  1     0
		 / \   / \
		1   0 1   1

	*/
	firstRoot := createNode(true)
	firstRoot.Left = createNode(true)
	firstRoot.Right = createNode(false)
	firstRoot.Left.Left = createNode(true)
	firstRoot.Left.Right = createNode(false)
	firstRoot.Right.Left = createNode(true)
	firstRoot.Right.Right = createNode(true)
	fmt.Println(unrollGarland(firstRoot))

	/*
				1
			   /  \
			  1     0
			 / \    / \
			1   0  1   1
		   / \ / \ / \ / \
		   1 0 0 1 1 1 1 1
	*/
	firstRoot.Left.Left.Left = createNode(true)
	firstRoot.Left.Left.Right = createNode(false)
	firstRoot.Left.Right.Left = createNode(false)
	firstRoot.Left.Right.Right = createNode(true)

	firstRoot.Right.Left.Left = createNode(true)
	firstRoot.Right.Left.Right = createNode(true)
	firstRoot.Right.Right.Left = createNode(true)
	firstRoot.Right.Right.Right = createNode(true)
	fmt.Println(unrollGarland(firstRoot))
	/*
				1
			   /  \
			  1     0
			 / \    / \
			1   0  1   1
		   / \ / \ / \ / \
		   1 0 0 1 1 1 1 1
		                / \
						0  1
	*/
	firstRoot.Right.Right.Right.Right = createNode(true)
	firstRoot.Right.Right.Right.Left = createNode(false)
	fmt.Println(unrollGarland(firstRoot))
	/*
				1
			   /  \
			  1     0
			 / \    / \
			1   0  1   1
		   / \ / \ / \ / \
		   1 0 0 1 1 1 1 1
		    / \         / \
		    1  0	    0  1
	*/
	firstRoot.Left.Left.Right.Left = createNode(true)
	firstRoot.Left.Left.Right.Right = createNode(false)
	fmt.Println(unrollGarland(firstRoot))
	/*
		1
	*/
	secondTree := createNode(true)
	fmt.Println(unrollGarland(secondTree))
	/*empty*/
	fmt.Println(unrollGarland(nil))
}

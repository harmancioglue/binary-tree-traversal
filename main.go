package main

import "fmt"

func main() {

	/*
			1
		   / \
		  2   3
		 / \
	    4   5

	*/

	root := &TreeNode{
		val:   1,
		left:  &TreeNode{
			val:   2,
			left:  &TreeNode{
				val:   4,
				left:  nil,
				right: nil,
			},
			right: &TreeNode{
				val:   5,
				left:  nil,
				right: nil,
			},
		},
		right: &TreeNode{
			val:   3,
			left:  nil,
			right: nil,
		},
	}

	output := preorderTraversal(root)
	fmt.Println(output) //[1 2 4 5 3]

	output = inOrderTraversal(root)
	fmt.Println(output) //[4 2 5 1 3]

	output = postorderTraversal(root)
	fmt.Println(output) //[4 5 2 3 1]

	levelOutput := levelOrderTraversal(root)
	fmt.Println(levelOutput) //[[1] [2 3] [4 5]]
}

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorder(root, &result)
	return result
}

func preorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.val)
	preorder(node.left, result)
	preorder(node.right, result)
}

func inOrderTraversal(root *TreeNode) []int {
	result := []int{}
	inorder(root, &result)
	return result
}
func inorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inorder(node.left, result)
	*result = append(*result, node.val)
	inorder(node.right, result)
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorder(root, &result)
	return result
}

func postorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	postorder(node.left, result)
	postorder(node.right, result)
	*result = append(*result, node.val)
}

func levelOrderTraversal(root *TreeNode) [][]int {
	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, node.val)

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

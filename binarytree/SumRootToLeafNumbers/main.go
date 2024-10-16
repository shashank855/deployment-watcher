package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructTree(index int, arr []int) *TreeNode {
	if index < len(arr) {
		TreeNode := new(TreeNode)
		TreeNode.Val = arr[index]
		TreeNode.Left = constructTree((2*index)+1, arr)
		TreeNode.Right = constructTree((2*index)+2, arr)
		return TreeNode
	} else {
		return nil
	}
}

// func preordertraversetree(root *TreeNode) {
// 	if root != nil {
// 		fmt.Println(root.Val)
// 		preordertraversetree(root.Left)
// 		preordertraversetree(root.Right)
// 	} else {
// 		return
// 	}
// }

func sumNumbers(root *TreeNode, parentNumber int) int {
	if root == nil {
		return 0
	}
	var sumleft, sumRight int

	if root.Left == nil && root.Right == nil {
		return root.Val + 10*parentNumber
	}
	if root.Left != nil {

		sumleft = sumNumbers(root.Left, 10*parentNumber+root.Val)
	}
	if root.Right != nil {

		sumRight = sumNumbers(root.Right, 10*parentNumber+root.Val)
	}
	return sumleft + sumRight
}

func main() {
	arr := []int{1, 2, 3}
	root := constructTree(0, arr)
	// preordertraversetree(root)
	fmt.Println(sumNumbers(root, 0))

}

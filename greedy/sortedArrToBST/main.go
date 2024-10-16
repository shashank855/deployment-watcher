package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func traverse(node *node) {
	if node == nil {
		return
	}
	traverse(node.left)
	fmt.Println(node.data)
	traverse(node.right)
}

func convertToBST(arr []int) *node {
	var index int
	if len(arr) > 0 {
		index = len(arr) / 2
		node := new(node)
		node.data = arr[index]
		leftSubArr := arr[:index]
		rightSubArr := arr[index+1:]
		node.left = convertToBST(leftSubArr)
		node.right = convertToBST(rightSubArr)
		return node
	}
	return nil
}

func main() {
	arr := []int{-10, -5, 0, 5, 8}
	root := convertToBST(arr)
	traverse(root)

}

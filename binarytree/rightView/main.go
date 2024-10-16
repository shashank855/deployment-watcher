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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	outputArr := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// outputArr = append(outputArr, root.Val)
	for len(queue) != 0 {
		outputArr = append(outputArr, queue[len(queue)-1].Val)
		initLenght := len(queue)
		for _, val := range queue {
			if val.Left != nil {
				queue = append(queue, val.Left)
			}
			if val.Right != nil {
				queue = append(queue, val.Right)
			}
		}
		queue = queue[initLenght:]
	}
	return outputArr
}

func main() {
	arr := []int{5, 6, 3, 2, 4, 9, 8}
	root := constructTree(0, arr)
	fmt.Println(rightSideView(root))

}

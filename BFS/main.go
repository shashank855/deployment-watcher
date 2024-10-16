package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func createNode(val int) *Node {
	node := new(Node)
	node.val = val
	node.left = nil
	node.right = nil
	return node
}

//[TODO]
//Traverse tree using BFS
func traverseTree(root *Node) {
	var queue []*Node
	queue = append(queue, root)

	for len(queue) != 0 {
		ptr := queue[0]
		fmt.Print("\nPopping:", queue[0].val)
		if ptr.left != nil {
			queue = append(queue, ptr.left)
		}
		if ptr.right != nil {
			queue = append(queue, ptr.right)
		}
		queue = queue[1:]
	}
}

func createBinarytree(arr []int, index int) *Node {
	if index < len(arr) {
		root := createNode(arr[index])
		root.left = createBinarytree(arr, (2*index + 1))
		root.right = createBinarytree(arr, (2*index + 2))
		return root
	}
	return nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := createBinarytree(arr, 0)
	traverseTree(root)
}

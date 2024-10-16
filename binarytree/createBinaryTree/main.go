package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

func constructTree(index int, arr []int) *node {
	if index < len(arr) {
		node := new(node)
		node.data = arr[index]
		node.left = constructTree((2*index)+1, arr)
		node.right = constructTree((2*index)+2, arr)
		return node
	} else {
		return nil
	}
}

func preordertraversetree(root *node) {
	if root != nil {
		fmt.Println(root.data)
		preordertraversetree(root.left)
		preordertraversetree(root.right)
	} else {
		return
	}
}

func levelordertraverseTree(root *node) {
	tmp_arr := []*node{}
	tmp_arr = append(tmp_arr, root)
	for len(tmp_arr) != 0 {
		if tmp_arr[0].left != nil {
			tmp_arr = append(tmp_arr, tmp_arr[0].left)
		}
		if tmp_arr[0].right != nil {
			tmp_arr = append(tmp_arr, tmp_arr[0].right)
		}
		fmt.Println(tmp_arr[0].data)
		tmp_arr = tmp_arr[1:]
	}
}

func main() {
	arr := []int{5, 6, 3, 2, 4, 9, 8}
	root := constructTree(0, arr)
	fmt.Println("**Printing level order**")
	levelordertraverseTree(root)
	fmt.Println("**Printing preorder**")
	preordertraversetree(root)

}

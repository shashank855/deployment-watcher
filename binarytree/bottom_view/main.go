package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

var my_map = make(map[int]int)

func preordertraversetree(root *node) {
	if root != nil {
		fmt.Println(root.data)
		preordertraversetree(root.left)
		preordertraversetree(root.right)
	} else {
		return
	}
}

func construct_Tree(index int, arr []int) *node {
	if index >= len(arr) {
		return nil
	}

	left_child_index := 2*index + 1
	right_child_index := 2*index + 2
	node := new(node)
	node.data = arr[index]
	node.left = construct_Tree(left_child_index, arr)
	node.right = construct_Tree(right_child_index, arr)

	return node
}

func botton_view(horizontal_distance int, root *node) {
	my_map[horizontal_distance] = root.data
	if root.left != nil {
		botton_view(horizontal_distance-1, root.left)
	}
	if root.right != nil {
		botton_view(horizontal_distance+1, root.right)
	}
	fmt.Println(my_map)

}

func main() {
	arr := []int{2, 4, 1, 5, 7, 29, 6, 8, 11}

	root_node := construct_Tree(0, arr)
	preordertraversetree(root_node)
	botton_view(0, root_node)

}

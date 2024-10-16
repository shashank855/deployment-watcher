package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	i := 1
	j := len(arr) - 1
	for j >= i {
		if arr[i] > pivot {
			if arr[j] < arr[i] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
				j--
			} else {
				j--
			}
		} else {
			i++
		}
	}
	tmp := arr[0]
	arr[0] = arr[j]
	arr[j] = tmp
	if j > 0 {
		quickSort(arr[:j])
		quickSort(arr[j+1:])
	}
}

type node struct {
	data  int
	left  *node
	right *node
}

func constructBST(arr []int, start int, end int) *node {

	if start > end {
		return nil
	}
	mid := (end + start) / 2
	node := new(node)
	node.data = arr[mid]
	node.left = constructBST(arr, start, mid-1)
	node.right = constructBST(arr, mid+1, end)
	return node
}

func traverseInorder(head *node) {
	if head == nil {
		return
	}
	traverseInorder(head.left)
	fmt.Println(head.data)
	traverseInorder(head.right)
}

func main() {
	arr := []int{4, 1, 8, 6, 2}
	//in order to convert normal array to BST first sort the array in ascending order
	quickSort(arr)
	fmt.Println(arr)
	head := constructBST(arr, 0, len(arr)-1)
	traverseInorder(head)
}

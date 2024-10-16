package main

import "fmt"

func swapData(i1 int, i2 int, arr []int) {
	tmp_val := arr[i1]
	arr[i1] = arr[i2]
	arr[i2] = tmp_val
}

func getGreaterValueIndex(left_index int, right_index int, arr []int) int {
	if arr[left_index] > arr[right_index] {
		return left_index
	} else {
		return right_index
	}
}

func heapify(arr []int, index int) {
	if index >= len(arr)-1 {
		return
	}

	left_child_index := 2*index + 1
	right_child_index := 2*index + 2

	gater_value_index := getGreaterValueIndex(left_child_index, right_child_index, arr)

	if arr[index] < arr[gater_value_index] {
		swapData(index, gater_value_index, arr)
		heapify(arr, gater_value_index)
	} else {
		return
	}
}

func createMaxHeap(arr []int) {
	last_non_leaf_index := len(arr)/2 - 1
	for i := last_non_leaf_index; i >= 0; i-- {
		heapify(arr, i)
	}
}

func main() {
	arr := []int{1, 7, 9, 6, 4, 5, 8}
	createMaxHeap(arr)
	fmt.Println(arr)
}

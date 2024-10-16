package main

import "fmt"

type node struct {
	sum       int
	left_val  int
	right_val int
}

var node_arr []node

func swap(index1 int, index2 int) {
	tmp_val := node_arr[index1]
	node_arr[index1] = node_arr[index2]
	node_arr[index2] = tmp_val
}

func heapify(index int) {
	if index == 0 {
		return
	}
	var parent_index int
	if index%2 == 0 {
		parent_index = (index - 2) / 2
	} else {
		parent_index = (index - 1) / 2
	}
	if node_arr[parent_index].sum >= node_arr[index].sum {
		swap(parent_index, index)
		heapify(parent_index)
	}

}

func insertIntoMinHeap(item *node) {
	if len(node_arr) == 0 {
		node_arr = append(node_arr, *item)
	} else {
		node_arr = append(node_arr, *item)
		heapify(len(node_arr) - 1)
	}
}

func getSmallerValueIndex(left_index int, right_index int) int {
	if node_arr[left_index].sum < node_arr[right_index].sum {
		return left_index
	} else {
		return right_index
	}
}

func adjustElements(index int) {
	if index >= (len(node_arr)/2)-1 {
		return
	} else {
		left_child_index := 2*index + 1
		right_child_index := 2*index + 2

		smaller_value_index := getSmallerValueIndex(left_child_index, right_child_index)
		swap(index, smaller_value_index)
		adjustElements(smaller_value_index)
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var result_arr [][]int
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			tmp_node := new(node)
			tmp_node.sum = nums1[i] + nums2[j]
			tmp_node.left_val = nums1[i]
			tmp_node.right_val = nums2[j]
			insertIntoMinHeap(tmp_node)
		}
	}
	for i := 0; i < k; i++ {
		result_arr = append(result_arr, []int{node_arr[0].left_val, node_arr[0].right_val})
		swap(0, len(node_arr)-1)
		node_arr = node_arr[:len(node_arr)-1]
		adjustElements(0)
	}
	return result_arr
}

func main() {
	nums1 := []int{1, 2, 4, 5, 6}
	nums2 := []int{3, 5, 7, 9}
	k := 3
	result := kSmallestPairs(nums1, nums2, k)
	fmt.Println(result)
}

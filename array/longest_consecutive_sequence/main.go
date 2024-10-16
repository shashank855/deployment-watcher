package main

import "fmt"

var my_map map[int]bool

var count int

func get_count(val int) int {
	var tmp_count int
	for my_map[val] {
		tmp_count = tmp_count + 1
		val = val + 1
	}

	if count > tmp_count {
		return count
	} else {
		return tmp_count
	}
}

func longestConsecutive(nums []int) int {
	// Initialize the map in the main function
	my_map = make(map[int]bool)
	count = 0
	for _, val := range nums {
		my_map[val] = true
	}

	for _, val := range nums {
		if my_map[val-1] == false {
			count = get_count(val)
		}
	}
	return count
}

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(arr))
}

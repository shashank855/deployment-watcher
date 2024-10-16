package main

import "fmt"

func max(arg1, arg2 int) int {
	if arg1 > arg2 {
		return arg1
	}
	return arg2
}

func rob(nums []int, start, end int) int {

	if end == start {
		return nums[start]
	}

	if end-start == 1 {
		return max(nums[end], nums[start])
	}

	if v, ok := cache[start]; ok {
		return v
	}

	answerWhenPicked := nums[start] + rob(nums, start+2, end)
	answerWhenNotPicked := rob(nums, start+1, end)
	answer := max(answerWhenPicked, answerWhenNotPicked)
	cache[start] = answer
	return answer
}

var cache map[int]int

func main() {
	arr := []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}
	//arr := []int{2, 7}
	cache = make(map[int]int)
	size := len(arr) - 1
	result := rob(arr, 0, size)
	fmt.Println(result)
}

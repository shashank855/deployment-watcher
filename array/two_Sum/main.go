package main

import "fmt"

func twoSum(nums []int, target int) []int {
	my_map := make(map[int]int)
	var indexArr []int
	for key, val := range nums {
		numbertoFindInMap := target - val
		_, exists := my_map[numbertoFindInMap]
		if exists {
			indexArr = append(indexArr, key, my_map[numbertoFindInMap])
		} else {
			my_map[val] = key
		}
	}
	return indexArr
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	outPut := twoSum(nums, target)
	fmt.Println(outPut)
}

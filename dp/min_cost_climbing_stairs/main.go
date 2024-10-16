package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(arr []int, cost int) int {
	if len(arr) == 0 {
		return cost
	}
	if len(arr) == 1 {
		return cost
	}
	return min(minCostClimbingStairs(arr[1:], arr[0]+cost), minCostClimbingStairs(arr[2:], arr[0]+cost))
}

func main() {
	arr := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	cost_from_index_zero := minCostClimbingStairs(arr[0:], arr[0])
	cost_from_index_one := minCostClimbingStairs(arr[1:], arr[1])
	fmt.Println(min(cost_from_index_one, cost_from_index_zero))
}

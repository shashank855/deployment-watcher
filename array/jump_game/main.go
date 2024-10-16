package main

import "fmt"

var global_result bool

func canJump(arr []int) bool {
	if global_result {
		return global_result
	}

	if len(arr) == 1 {
		global_result = true
		return true
	}
	if arr[0] != 0 {
		for i := 1; i <= arr[0]; i++ {
			result := canJump(arr[i:])
			if result {
				global_result = true
				return true
			}
		}
	}
	return false
}

func main() {
	arr := []int{3, 2, 1, 0, 4}
	// canJump(arr)
	fmt.Println(canJump(arr))
}

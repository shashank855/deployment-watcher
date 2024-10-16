package main

import "fmt"

var min_jump int = 10000

func canJump(arr []int, jumps int) int {

	if len(arr) == 1 {
		return jumps
	}
	if arr[0] != 0 {
		for i := 1; i <= arr[0]; i++ {
			if i < len(arr) {
				result := canJump(arr[i:], jumps+1)
				if result < min_jump {
					min_jump = result
				}
			}
		}
	} else {
		return 10000
	}
	return min_jump
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(canJump(arr, 0))

}

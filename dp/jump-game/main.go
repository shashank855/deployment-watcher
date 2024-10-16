package main

import "fmt"

var result bool

func newcanjump(index int, arr []int) bool {
	if index == len(arr)-1 {
		return true
	}

	for i := 1; i <= arr[index]; i++ {
		if newcanjump(index+i, arr) {
			result = true
			break
		}
	}
	return result
}

func canJump(arr []int) bool {
	current_index := 0
	return newcanjump(current_index, arr)
}

func main() {
	arr := []int{3, 2, 1, 0, 4}
	result = canJump(arr)
	fmt.Println(result)
}

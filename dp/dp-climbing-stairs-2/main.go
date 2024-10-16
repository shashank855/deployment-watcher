package main

import "fmt"

var status bool

func canReach(arr []int, index int) bool {
	if (len(arr) - 1) <= index {
		fmt.Println("true hua")
		return true
	}
	jump := arr[index]
	if jump != 0 {
		for i := 1; i <= jump; i++ {
			canReach(arr, (index + i))
			fmt.Println("----iteration:", i)
		}
	}
	return false
}

func main() {
	arr := []int{1, 1, 1}
	if canReach(arr, 0) == true {
		fmt.Println("can reach")
	} else {
		fmt.Println("can not reach")
	}
}

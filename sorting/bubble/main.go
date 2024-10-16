package main

import "fmt"

var count int

func buubleSort(arr []int) []int {
	for true {
		for i := 0; i < (len(arr) - 1); i++ {
			if arr[i] > arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				count = count + 1
			}
		}
		if count == 0 {
			break
		}
		count = 0
	}
	return arr
}

func main() {
	arr := []int{7, 6, 5, 4, 3, 21}
	sortedArr := buubleSort(arr)
	for _, val := range sortedArr {
		fmt.Println(val)
	}
}

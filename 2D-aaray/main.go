package main

import (
	"fmt"
)

func main() {
	fmt.Println("Add values in array 1")
	array_1 := [3][3]int
	fmt.Println("Enter data")
	var i, j, num int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Scanf("%d", &num)
			array_1[i][j] = num

		}
	}

	fmt.Println("Elements in array-1")
	for i = 0; i < 3; i++ {
		var count int
		count = 0
		for j = 0; j < 3; j++ {
			fmt.Printf("%d  ", array_1[i][j])
			count++
			if count == 3 {
				fmt.Printf("\n")
				count = 0
			}
		}
	}

	fmt.Println("Add values in array 2")
	array_2 := [3][3]int{}
	fmt.Println("Enter data")
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Scanf("%d", &num)
			array_2[i][j] = num

		}
	}

	fmt.Println("Elements in array-2")
	for i = 0; i < 3; i++ {
		var count int
		count = 0
		for j = 0; j < 3; j++ {
			fmt.Printf("%d  ", array_2[i][j])
			count++
			if count == 3 {
				fmt.Printf("\n")
				count = 0
			}
		}
	}

	fmt.Println("Their Addition:")

	for i = 0; i < 3; i++ {
		var count, val int
		count = 0
		for j = 0; j < 3; j++ {
			val = array_1[i][j] + array_2[i][j]
			fmt.Printf("%d  ", val)
			count++
			if count == 3 {
				fmt.Printf("\n")
				count = 0
			}
		}
	}
}

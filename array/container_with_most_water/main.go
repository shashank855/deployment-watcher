package main

import "fmt"

func maxArea(height []int) int {
	var breath, length int

	finalArea := 0
	i := 0
	j := len(height) - 1
	for i < j {
		if height[i] > height[j] {
			length = height[j]
		} else {
			length = height[i]
		}
		breath = j - i
		currentArea := length * breath
		if currentArea > finalArea {
			finalArea = currentArea
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return finalArea
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(arr))

}

package main

import "fmt"

func missingNumber(nums []int) int {
	//[3,0,1] Consider this is artihmatic progression of 0,1,2,3
	// AP Sum = sum of elements present + missing element
	var sum int
	apSum := (len(nums) + 1) * (len(nums)) / 2
	fmt.Println(apSum)
	for _, val := range nums {
		sum = sum + val
	}
	fmt.Println(sum)
	return apSum - sum

}

func main() {
	arr := []int{0, 2, 3, 4}
	ele := missingNumber(arr)
	fmt.Println(ele)
}

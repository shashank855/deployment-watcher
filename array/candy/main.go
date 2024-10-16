package main

import "fmt"

func candy(ratings []int) {
	var number_of_candies []int
	var total_candy int
	for i := 0; i <= (len(ratings) - 1); i++ {
		if i == 0 {
			number_of_candies = append(number_of_candies, 1)
		} else {
			if ratings[i] > ratings[i-1] {
				number_of_candies = append(number_of_candies, number_of_candies[i-1]+1)
			} else {
				number_of_candies = append(number_of_candies, 1)
			}
		}
	}
	fmt.Println("left2right:", number_of_candies)

	for i := (len(ratings) - 2); i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if number_of_candies[i] <= number_of_candies[i+1] {
				number_of_candies[i] = number_of_candies[i+1] + 1
			}
		}
	}

	fmt.Println("right2left:", number_of_candies)

	for _, val := range number_of_candies {
		total_candy = total_candy + val
	}
	fmt.Println(total_candy)
}

func main() {
	//arr := []int{1, 6, 10, 8, 7, 3, 2}
	arr := []int{1, 3, 5, 4, 2}
	candy(arr)
}

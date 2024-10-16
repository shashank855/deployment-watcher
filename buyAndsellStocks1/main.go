package main

import (
	"fmt"
)

// func maxProfit(prices []int) int {
// 	valey := 1000
// 	profit := 0
// 	for i := 0; i < (len(prices) - 1); i++ {
// 		if prices[i] < valey {
// 			valey = prices[i]
// 		} else if profit < (prices[i+1] - valey) {
// 			profit = prices[i+1] - valey
// 		}

// 	}
// 	return (profit)
// }

func maxProfit(prices []int) int {
	var profit, curreentProfit int = 0, 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit = prices[j] - prices[i]
			if profit > curreentProfit {
				curreentProfit = profit
			}
		}
	}
	return curreentProfit
}

func main() {
	//{3, 2, 6, 0, 3}
	//{2, 1, 2, 0, 1}
	//{7, 1, 5, 3, 6, 4}
	//{1, 2, 4}
	//{2, 1, 2, 1, 0, 1, 2}
	arr := []int{1, 2, 4}
	max := maxProfit(arr)
	fmt.Println("MAx is:", max)
}

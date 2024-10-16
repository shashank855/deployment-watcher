package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func calculate_coins(coins []int, amount int, steps int) int {

	if len(coins) == 1 && coins[0] < amount {
		return 100000
	}
	if amount == 0 {
		return steps
	}
	if amount < 0 {
		return 100000
	}
	if len(coins) == 0 {
		return 100000
	}
	fmt.Printf("coins arr:%d \n Amount:%d \n steps:%d \n ======\n", coins, amount, steps)
	mincoins := min(calculate_coins(coins, amount-coins[0], steps+1), calculate_coins(coins[1:], amount, steps))
	//time.Sleep(time.Second * 5)
	return mincoins

}

func coinChange(coins []int, amount int) int {
	total_coins := calculate_coins(coins, amount, 0)
	if total_coins == 100000 {
		return -1
	} else {
		return total_coins
	}
}

func main() {
	arr := []int{1, 2, 5}
	fmt.Println(coinChange(arr, 10))
}

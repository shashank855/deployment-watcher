package main

import "fmt"

var dict = make(map[int]int)

func climbStairs(n int) int {

	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if dict[n] == 0 {
		dict[n] = climbStairs(n-1) + climbStairs(n-2)
	}
	return dict[n]

}

func main() {
	ways := climbStairs(5)
	fmt.Println(ways)
}

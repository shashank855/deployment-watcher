package main

import "fmt"

func solve(n int, s int, d int, h int) {
	if n == 1 {
		fmt.Printf("Moving plate- %d from pole-%d to pole-%d", n, s, d)
		fmt.Println("")
		return
	}

	solve(n-1, s, h, d)
	fmt.Printf("Moving plate-%d from pole-%d to pole-%d", n, s, d)
	fmt.Println("")
	solve(n-1, h, d, s)
}
func main() {
	n := 3
	s := 1 //source
	h := 2 //helper
	d := 3 //destination
	solve(n, s, d, h)
}

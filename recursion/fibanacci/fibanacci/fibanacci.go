package fibanacci

var my_map = map[int]int{}

func Fib(n int) int {
	if n <= 1 {
		my_map[n] = n
		return n
	}
	if my_map[n] == 0 {
		my_map[n] = Fib(n-1) + Fib(n-2)
		return my_map[n]
	} else {
		return my_map[n]
	}
}

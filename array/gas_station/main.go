package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	diff_arr := make([]int, len(gas))
	var possible_index []int
	var final_index, final_sum int
	for i := 0; i <= len(gas)-1; i++ {
		diff_arr[i] = gas[i] - cost[i]
	}
	for key, val := range diff_arr {
		if val > 0 {
			possible_index = append(possible_index, key)
		}
		final_sum = final_sum + val
	}
	fmt.Println(diff_arr)
	fmt.Println("possible index:", possible_index)
	sum_from_starting_index := 0
outerloop:
	for j := 0; j <= (len(possible_index) - 1); j++ {
		sum_from_starting_index = 0

	innerloop:
		for i := 0; i <= len(gas)-1; i++ {

			sum_from_starting_index = sum_from_starting_index + diff_arr[(possible_index[j]+i)%len(gas)]
			fmt.Println("Index:", i)
			fmt.Println("Sum:", sum_from_starting_index)
			if sum_from_starting_index < 0 {
				break innerloop
			}
		}

		if sum_from_starting_index >= 0 {
			fmt.Println("Enterred here")
			final_index = possible_index[j]
			break outerloop
		}
		final_index = possible_index[j]

	}
	if final_sum >= 0 {
		return final_index
	} else {
		return -1
	}
}

func main() {
	gas := []int{5, 1, 2, 3, 4}
	cost := []int{4, 4, 1, 5, 1}
	fmt.Println(canCompleteCircuit(gas, cost))
}

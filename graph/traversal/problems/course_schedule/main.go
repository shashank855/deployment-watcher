package main

import "fmt"

var count int = 0
var cycle_present bool

func topological_sort(my_map map[int][]int, visited_map map[int]bool, current_node int, parent_node int) {

	for _, val := range my_map[current_node] {
		if visited_map[val] == false && len(my_map[current_node]) > 0 {
			if val == parent_node { // cycle detected
				cycle_present = true
				break
			}
			topological_sort(my_map, visited_map, val, current_node)
		}
	}
	if visited_map[current_node] == false {
		count = count + 1
	}
	visited_map[current_node] = true
}

func main() {

	// my_map := map[int][]int{
	// 	5: {0, 2},
	// 	2: {3},
	// 	3: {1},
	// 	4: {0, 1},
	// 	1: {},
	// 	0: {},
	// }

	my_map := make(map[int][]int)
	my_arr := [][]int{
		// {1, 0},
		{0, 1},
	}
	for _, item := range my_arr {
		my_map[item[0]] = append(my_map[item[0]], item[1])
		my_map[item[1]] = append(my_map[item[1]])

	}
	fmt.Println(my_map)
	visited_map := make(map[int]bool)
	for item := range my_map {
		if visited_map[item] == false {
			topological_sort(my_map, visited_map, item, -1)
		}
	}
	fmt.Println(count)
	if count == len(my_map) && cycle_present == false {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

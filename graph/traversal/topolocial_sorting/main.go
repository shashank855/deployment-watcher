package main

import "fmt"

var visted_map = make(map[int]int)

func topological_sort(key int, my_map map[int][]int) {
	if len(my_map[key]) == 0 && visted_map[key] == 0 {
		visted_map[key] = 1
		fmt.Println(key)
		return
	}
	for _, val := range my_map[key] {
		topological_sort(val, my_map)
	}
	if visted_map[key] == 0 {
		fmt.Println(key)
		visted_map[key] = 1
	}
}

func main() {
	my_map := map[int][]int{
		5: {0, 2},
		2: {3},
		3: {1},
		4: {0, 1},
		1: {},
		0: {5},
	}
	for key, _ := range my_map {
		if visted_map[key] == 0 {
			topological_sort(key, my_map)
		}
	}
}

package main

import "fmt"

func find_connectivity(my_map map[int][]int, source int, destination int) bool {
	visited_map := make(map[int]int)
	var queue []int
	queue = append(queue, source)
	for len(queue) != 0 {
		if queue[0] == destination {
			return true
		}
		if visited_map[queue[0]] == 0 {
			for _, val := range my_map[queue[0]] {
				queue = append(queue, val)
			}
		}
		if visited_map[queue[0]] == 0 {
			visited_map[queue[0]] = 1
		}
		queue = queue[1:]
	}
	return false
}

func main() {
	//[[0,1],[0,2],[3,5],[5,4],[4,3]]
	my_map := make(map[int][]int)

	my_arr := [][]int{
		{0, 1},
		{0, 2},
		{3, 5},
		{5, 4},
		{4, 3},
	}
	for _, val := range my_arr {
		my_map[val[0]] = append(my_map[val[0]], val[1])
		my_map[val[1]] = append(my_map[val[1]], val[0])
	}
	fmt.Println(my_map)
	fmt.Println(find_connectivity(my_map, 5, 3))
}

package main

import "fmt"

func set_Rev_colour(colour string) string {
	if colour == "red" {
		return "blue"
	} else if colour == "blue" {
		return "red"
	} else {
		return "red"
	}
}

type visit struct {
	visited bool
	colour  string
}

var visited_nodes = map[int]visit{}

func is_2mColorable(my_map map[int][]int, root int) bool {

	var queue []int
	queue = append(queue, root)

	for len(queue) != 0 {
		if visited_nodes[queue[0]].visited == false && visited_nodes[queue[0]].colour == "nil" {
			visited_nodes[queue[0]] = visit{
				visited: true,
				colour:  set_Rev_colour(""),
			}
		}
		current_colour := visited_nodes[queue[0]].colour
		for _, val := range my_map[queue[0]] {
			if current_colour == visited_nodes[val].colour {
				return false
			}
			if visited_nodes[val].visited == false {
				queue = append(queue, val)
				visited_nodes[val] = visit{
					visited: true,
					colour:  set_Rev_colour(current_colour),
				}
			}
		}
		fmt.Println("current_colour", current_colour)
		fmt.Println("queue:", queue)
		fmt.Println("visit_map", visited_nodes)
		queue = queue[1:]
	}

	return true
}

func main() {
	//This is implemented using 2-m colouring algorithm
	//[[1,2,3],[0,2],[0,1,3],[0,2]]

	my_arr := [][]int{
		{1, 2},
		{0, 3, 2},
		{0, 3},
		{2, 1},
		{8, 7},
	}
	my_map := make(map[int][]int)
	for key, item := range my_arr {
		for _, val := range item {
			my_map[key] = append(my_map[key], val)
		}
	}
	fmt.Println(my_map)
	for key := range my_map {
		visited_nodes[key] = visit{
			visited: false,
			colour:  "nil",
		}
	}
	var iscolourable bool = true

	for key := range visited_nodes {
		if visited_nodes[key].visited == false {
			iscolourable = is_2mColorable(my_map, key) && iscolourable
		}
	}
	fmt.Println(iscolourable)
}

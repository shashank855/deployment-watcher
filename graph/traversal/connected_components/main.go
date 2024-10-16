package main

import (
	"fmt"
)

var number_of_components int = 0
var larget_component int = 0
var visited_map = make(map[string]int)

func traverse_graph(my_map map[string][]string, key string) {
	var stack []string
	largest_piece := 0
	stack = append(stack, key)
	fmt.Println(stack)
	for len(stack) != 0 {
		current_val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, val := range my_map[current_val] {
			if visited_map[val] == 0 {
				stack = append(stack, val)
				visited_map[val] = 1
			}
		}
		largest_piece = largest_piece + 1
	}
	if larget_component < largest_piece {
		larget_component = largest_piece
	}
	number_of_components = number_of_components + 1
}

func get_connected_components(my_map map[string][]string) {
	for key := range my_map {
		visited_map[key] = 0
	}
	for key := range my_map {
		if visited_map[key] == 0 {
			traverse_graph(my_map, key)
		}
	}
}

func main() {
	my_map := map[string][]string{
		"a": {},
		"b": {"c"},
		"c": {"b"},
		"d": {"e"},
		"e": {"d", "f", "g", "h"},
		"f": {"e"},
		"g": {"e"},
		"h": {"e"},
	}
	get_connected_components(my_map)
	fmt.Println(number_of_components)
	fmt.Println(larget_component)
}

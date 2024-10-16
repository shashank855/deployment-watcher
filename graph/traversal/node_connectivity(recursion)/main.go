package main

import "fmt"

var check bool

func is_neighbour(source string, neighbour string, my_map map[string][]string) bool {
	for _, val := range my_map[source] {
		if val == neighbour {
			return true
		}
	}
	return false
}

func can_traverse(source string, destination string, my_map map[string][]string) bool {
	if source != "" && destination != "" {
		if is_neighbour(source, destination, my_map) {
			return true
		}
		for _, val := range my_map[source] {
			check = check || can_traverse(val, destination, my_map)
		}
	}
	return check
}

func main() {
	my_map := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {},
		"e": {"b"},
		"f": {"d"},
	}
	fmt.Println(can_traverse("a", "f", my_map))
}

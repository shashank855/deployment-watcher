package main

import "fmt"

func bfs_traverse(my_map map[string][]string, root string) {
	var queue []string
	visited_map := make(map[string]int)
	queue = append(queue, root)
	for len(queue) != 0 {
		if visited_map[queue[0]] == 0 { // marking the visited node other wise it will give >
			fmt.Println(queue[0]) // > redundant output
		}
		for _, val := range my_map[queue[0]] {
			if visited_map[queue[0]] == 0 {
				queue = append(queue, val)
			}
		}
		if visited_map[queue[0]] == 0 {
			visited_map[queue[0]] = 1
		}
		queue = queue[1:]
	}
}

func main() {
	my_map := map[string][]string{
		// "a": {"c", "b"},
		// "b": {"d"},
		// "c": {"e"},
		// "d": {},
		// "e": {"b"},
		// "f": {"d"},
		"b": {"c"},
		"c": {"b"},
	}
	bfs_traverse(my_map, "b")
}

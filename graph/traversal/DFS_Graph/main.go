package main

import "fmt"

func dfs_traverse(my_map map[string][]string, root string) {
	var stack []string
	visited_map := make(map[string]int)
	stack = append(stack, root)
	for len(stack) != 0 {
		if visited_map[stack[len(stack)-1]] == 0 { // marking the visited node other wise it will give >
			fmt.Println(stack[(len(stack) - 1)]) // > redundant output
		}
		tmp_val := stack[(len(stack) - 1)]
		stack = stack[:(len(stack) - 1)]
		for _, val := range my_map[tmp_val] {
			stack = append(stack, val)
		}
		if visited_map[tmp_val] == 0 {
			visited_map[tmp_val] = 1
		}
	}
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
	dfs_traverse(my_map, "a")

}

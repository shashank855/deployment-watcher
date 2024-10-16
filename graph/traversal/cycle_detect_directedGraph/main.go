package main

import "fmt"

var visit_map = make(map[int]int)
var dfs_node_instack = make(map[int]int)

func modified_dfs(my_map map[int][]int, key int) {
	dfs_node_instack[key] = 1
	if visit_map[key] == 0 {
		visit_map[key] = 1
		for _, val := range my_map[key] {
			if dfs_node_instack[val] == 1 {
				fmt.Println("cycle Detected")
				break
			}
			modified_dfs(my_map, val)
		}
		dfs_node_instack[key] = 0
	} else {
		dfs_node_instack[key] = 0
	}
}

func main() {
	/* - We will use DFS(it means we will use stack) for cycle detection in directed graph.
	- We will maintain the visted map and a map for DFS's nodes in recursion. The DFS nodes which
	will be in recursion will be marked as 1.
	- As soon the recursion folds we will mark dfs nodes out of recursion as 0 again
	Logic is as soon we get any node which is DFS's recursion map we will say there is cycle*/
	my_map := map[int][]int{
		5: {0, 2},
		2: {3},
		3: {1},
		4: {0, 1},
		1: {},
		0: {},
	}
	for key := range my_map {
		modified_dfs(my_map, key)
	}
}

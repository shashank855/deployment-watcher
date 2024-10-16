package main

import "fmt"

type node struct {
	prev *node
	data int
	next *node
}

func createCircularLinklist(num_players int) *node {
	var root, current_ptr *node
	for i := 1; i <= num_players; i++ {
		if i == 1 {
			root = new(node)
			root.data = i
			root.next = nil
			root.prev = nil
			current_ptr = root
		} else {
			tmp_ptr := current_ptr
			current_ptr.next = new(node)
			current_ptr = current_ptr.next
			current_ptr.data = i
			current_ptr.prev = tmp_ptr
			if i == num_players {
				current_ptr.next = root
				root.prev = current_ptr
			}
		}
	}
	return root
}

func eliminate(root *node, interval int) *node {
	cur_ptr := root
	for i := 0; i < (interval - 1); i++ {
		cur_ptr = cur_ptr.next
	}
	new_root := cur_ptr.next
	cur_ptr.next = nil
	cur_ptr = cur_ptr.prev
	cur_ptr.next = new_root
	new_root.prev = cur_ptr
	return new_root
}

func find_winner(num_players int, interval int) {
	root := createCircularLinklist(num_players)
	for i := 0; i < num_players; i++ {
		root = eliminate(root, interval)

	}
	fmt.Println(root)
}

func main() {
	find_winner(5, 2)
}

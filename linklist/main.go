package main

import (
	"fmt"
)

type Node struct {
	data int
	ptr  *Node
}

var head, tail *Node

func bindNode(tmp *Node) {
	if head == nil {
		head = tmp
		tail = tmp
	} else {
		tail.ptr = tmp
		tail = tmp
	}
}

func addNode(val int) *Node {
	node := new(Node)
	node.data = val
	node.ptr = nil
	return node
}

func createlist() {
	var val, num int
	fmt.Println("1.Add entry")
	fmt.Println("2.Done")
	fmt.Println("Enter Your Options:")
	fmt.Scanf("%d", &num)

	switch options := num; options {
	case 1:
		fmt.Println("Enter DATA")
		fmt.Scanf("%d", &val)
		current_node_adr := addNode(val)
		bindNode(current_node_adr)
		createlist()
	case 2:
		printList()
		return
	default:
		fmt.Print("This is not a valid option\n")
		createlist()
	}

}

func printList() {
	fmt.Println("Node: ", head)
	tmp := head
	for {
		if tmp.ptr == nil {
			fmt.Printf("%d\n", tmp.data)
			break
		} else {
			fmt.Printf("%d\n", tmp.data)
			tmp = tmp.ptr
		}

	}
}

func main() {
	var num int
	head = nil
	tail = nil
	fmt.Print("1. Create linklist\n")
	fmt.Println("2. Exit")
	fmt.Println("Enter Your Options:")
	fmt.Scanf("%d", &num)
	switch options := num; options {
	case 1:
		createlist()
	case 2:
		fmt.Print("Exiting...")
	default:
		fmt.Print("This is not a valid option")
	}
}

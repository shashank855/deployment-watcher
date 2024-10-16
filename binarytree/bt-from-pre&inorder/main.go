package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getIndex(value int, inorder []int) int {
	var index int
	for key, val := range inorder {
		if val == value {
			index = key
			break
		}
	}
	return index
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	node := new(TreeNode)

	if len(inorder) == 1 {
		node.Val = preorder[0]
		node.Left = nil
		node.Right = nil
		fmt.Println("inside if")
		fmt.Println("pre:", preorder)
		fmt.Println("in:", inorder)
		return node
	} else {
		fmt.Println("inside else")
		indexOfRootNode := getIndex(preorder[0], inorder)
		fmt.Println("indexofrootnode:", indexOfRootNode)
		fmt.Println("pre:", preorder)
		fmt.Println("in:", inorder)
		node.Val = preorder[0]
		preorder = preorder[1:]
		node.Left = buildTree(preorder, inorder[:indexOfRootNode])
		if len(preorder) > 1 {
			preorder = preorder[1:]
			node.Right = buildTree(preorder, inorder[indexOfRootNode+1:])
		}
	}
	return node
}

func traverseTree(root *TreeNode) {
	if root == nil {
		return
	} else {
		fmt.Println(root.Val)
		traverseTree(root.Left)
		traverseTree(root.Right)
	}
}

func main() {
	preorder := []int{1, 2}
	inorder := []int{1, 2}

	rootNode := buildTree(preorder, inorder)
	traverseTree(rootNode)

}

package main

import (
	"fmt"
)

type TreeNode struct {
	ID    int
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	node7 := &TreeNode{
		ID:    7,
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	node6 := &TreeNode{
		ID:    6,
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node5 := &TreeNode{
		ID:    5,
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node4 := &TreeNode{
		ID:    4,
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node3 := &TreeNode{
		ID:    3,
		Val:   3,
		Left:  node6,
		Right: node7,
	}
	node2 := &TreeNode{
		ID:    2,
		Val:   2,
		Left:  node4,
		Right: node5,
	}

	node1 := &TreeNode{
		ID:    1,
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	fmt.Println("先序遍历")
	PreOrder(node1)
	fmt.Println()
	fmt.Println("中序遍历")
	InOrder(node1)
	fmt.Println()
	fmt.Println("后序遍历")
	PostOrder(node1)
}

func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

func InOrder(root *TreeNode) {
	if root != nil {
		InOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		InOrder(root.Right)
	}
}

func PostOrder(root *TreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}

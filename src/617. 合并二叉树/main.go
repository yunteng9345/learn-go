package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// BFS
func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	merged := &TreeNode{Val: t1.Val + t2.Val}
	queue := []*TreeNode{merged}
	queue1 := []*TreeNode{t1}
	queue2 := []*TreeNode{t2}
	for len(queue1) > 0 && len(queue2) > 0 {
		node := queue[0]
		queue := queue[1:]
		node1 := queue1[0]
		queue1 := queue1[1:]
		node2 := queue2[0]
		queue2 := queue2[1:]
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right

	}

}

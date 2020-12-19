package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入： 1->2->3->4->5 和 k = 2
// 输出： 4

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	former, latter := head, head
	for k > 0 {
		former = former.Next
		k--
	}
	for former != nil {
		former = former.Next
		latter = latter.Next
	}
	return latter.Val
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	former, latter := head, head
	for i := 0; i < k; i++ {
		former = former.Next
	}
	for former != nil {
		former = former.Next
		latter = latter.Next
	}
	return latter
}

func main() {

}

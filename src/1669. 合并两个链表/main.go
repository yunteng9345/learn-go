package main

func main() {

}

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

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var preNode, postNode *ListNode
	p1 := list1
	i := 0
	for p1 != nil {
		i++
		preNode = p1
		if i == a {
			postNode = p1.Next
			p1.Next = list2
		} else {
			p1 = p1.Next
		}
	}
	i = a
	p1 = postNode
	for p1 != nil {
		if i == b {
			preNode.Next = p1.Next
			break
		} else {
			p1 = p1.Next
		}
		i++
	}
	return list1
}

func mergeInBetween1(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var preNode ,postNode *ListNode
	p1 := list1
	i := 0
	for p1 != nil {
		i++
		preNode = p1
		if i == a {
			postNode = p1.Next
			p1.Next = list2
		} else {
			p1 = p1.Next
		}
	}
	i = a
	p1 = postNode
	for p1 != nil {
		if i== b{
			preNode.Next = p1.Next
			break
		} else {
			p1 = p1.Next
		}
		i++
	}
	return list1
}

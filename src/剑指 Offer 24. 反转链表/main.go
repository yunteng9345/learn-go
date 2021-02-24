package main

import (
	"strconv"
	"strings"
)

func main() {
	reverseList2(CreateTestData("[1,2,3,4,5]"))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针
func reverseList1(head *ListNode) *ListNode {
	var cur *ListNode
	pre := head
	for pre != nil {
		t := pre.Next
		pre.Next = cur
		cur = pre
		pre = t
	}
	return cur
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

func CreateTestData(data string) *ListNode {
	if data == "[]" {
		return nil
	}
	data = string([]rune(data)[1 : len(data)-1])
	res := strings.Split(data, ",")
	length := len(res)
	listNode := make([]ListNode, length)
	headVal, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}
	listNode[0] = ListNode{headVal, nil}
	for i := 1; i < length; i++ {
		headVal, _ = strconv.Atoi(res[i])
		listNode[i] = ListNode{headVal, nil}
		listNode[i-1].Next = &listNode[i]
	}
	return &listNode[0]
}
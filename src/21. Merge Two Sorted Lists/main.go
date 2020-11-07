package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.
Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l1 := CreateTestData("[]")
	//l2 := CreateTestData("[]")
	//l3 := mergeTwoLists(l1, l2)
	//
	//Print(l3)
}

/**
* Definition for singly-linked list.
* type node struct {
*     Val int
*     Next *node
* }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}
	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}
	return result.Next
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

func Print(listNode *ListNode) {
	if listNode == nil {
		fmt.Println(nil)
	}
	var buffer strings.Builder
	buffer.WriteString("[")
	value := strconv.Itoa(listNode.Val)
	buffer.WriteString(value)
	temp := listNode.Next
	for temp != nil {
		buffer.WriteString(",")
		value = strconv.Itoa(temp.Val)
		buffer.WriteString(value)
		temp = temp.Next
	}
	buffer.WriteString("]")
	fmt.Println(buffer.String())
}

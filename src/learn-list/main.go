package main

import (
	"fmt"
)

type Node struct {
	element int
	next    *Node
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func LinkList() *List {
	list := new(List)
	tail := new(Node)
	list.tail = tail
	list.head = tail
	return list
}

func (list *List) append(e int) {
	node := new(Node)
	node.element = e
	list.head.next = node
	list.head = node
	list.len++
}

func (list *List) get(index int) int {

	var node *Node
	node = list.tail
	for i := 0; i <= index; i++ {
		node = node.next
		if node == nil {
			return 0
		}
	}
	return node.element
}
func (list *List) getNode(index int) *Node {
	var node *Node
	node = list.tail
	for i := 0; i <= index; i++ {
		node = node.next
	}
	return node
}

func (list *List) remove(index int) interface{} {
	var prev, current, next *Node
	prev = list.tail
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	current = prev.next
	next = current.next
	if list.len == index+1 {
		list.head = prev
		prev.next = nil
	} else {
		prev.next = next
		list.head = next
	}
	list.len--
	return current.element
}

func main() {
	list := LinkList()
	list.append(10)
	fmt.Println(list.get(0))
	a := list.remove(0)
	fmt.Println(a)
	list.append(20)
	list.append(11)
	list.append(114)
	list.append(115)
	list.append(181)
	list.append(4)
	b := list.remove(list.len - 1)
	fmt.Println(b)
	fmt.Println(list.get(0), list.len)
}

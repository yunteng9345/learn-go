package main

import "fmt"

func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	//http.ListenAndServe(":8080", nil)
	fmt.Println("hello World")

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
func deleteNode(node *ListNode) {

}

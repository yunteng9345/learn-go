package main

import "fmt"

/*
输入: s = "abcdefg", k = 2
输出: "cdefgab"
*/
func main() {
	s := "abcdefg"
	n := 2
	fmt.Println(s[n:] + s[:n])
	//fmt.Println(reverseLeftWords(s, n))
}
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

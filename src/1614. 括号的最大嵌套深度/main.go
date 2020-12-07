package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(()(()))"))
}

func maxDepth(s string) int {
	var pair, max int
	for _, v := range s {
		switch v {
		case '(':
			pair++
		case ')':
			if pair > max {
				max = pair
			}
			pair--
		}
	}
	return max
}

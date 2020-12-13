package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

func balancedStringSplit(s string) int {
	res := 1
	for i := range s {
		if i > 0 && s[i] != s[i-1] {
			res++
		}
	}
	fmt.Println("hello")
	return res
}

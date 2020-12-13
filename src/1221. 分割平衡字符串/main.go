package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLLLLRRRLR"))
}

var res int

func balancedStringSplit(s string) int {
	var index int
	for i := 0; i < len(s); i++ {
		temp := s[0]
		if temp != s[i] {
			index = 2*i - 1
			break
		}
	}
	if len(s) != 0 {
		res++
		balancedStringSplit(s[index+1:])
	}
	return res
}

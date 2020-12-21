package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}

// (()())   (())  (()(()))
// ()()()()(())
// ()() + () + ()()
func removeOuterParentheses(S string) string {
	flag := 0
	b := 1
	r := ""
	for i := range S {
		if S[i] == '(' {
			flag++
		} else {
			flag--
		}
		if flag == 0 {
			r += S[b:i]
			b = i + 2
		}
	}
	return r
}

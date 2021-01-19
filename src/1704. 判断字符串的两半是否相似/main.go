package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(halvesAreAlike("AbCdEfGh"))
}
func halvesAreAlike(s string) bool {
	a := "aeiouAEIOU"
	b := 0
	c := 0
	for i, v := range s {
		if i > len(s)/2 -1 && strings.ContainsAny(a, string(v)) {
			c++
		} else {
			if strings.ContainsAny(a, string(v)) {
				b++
			}
		}
	}
	return b == c
}

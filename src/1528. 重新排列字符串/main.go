package main

import (
	"fmt"
)

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}

func restoreString(s string, indices []int) string {
	var b [100]byte
	for i, v := range indices {
		b[v] = s[i]
	}
	return string(b[:len(s)])
}

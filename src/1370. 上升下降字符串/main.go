package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
}

// 桶排序
func sortString(s string) string {
	var hash [26]byte
	var b strings.Builder
	for _, c := range s {
		hash[c-'a']++
	}
	l := len(s)
	for l > 0 {
		for i := 0; i < 26; i++ {
			if hash[i] > 0 {
				b.WriteByte('a' + byte(i))
				hash[i]--
				l--
			}
		}
		for i := 25; i >= 0; i-- {
			if hash[i] > 0 {
				b.WriteByte('a' + byte(i))
				hash[i]--
				l--
			}
		}
	}
	return b.String()
}

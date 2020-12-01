package main

import (
	"fmt"
)

func main() {
	w1 := []string{"ab", "c"}
	w2 := []string{"a", "bc"}
	fmt.Println(arrayStringsAreEqual(w1, w2))

}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1 := ""
	w2 := ""
	for _, v := range word1 {
		w1 = w1 + v
	}
	for _, v := range word2 {
		w2 = w2 + v
	}
	return w1 == w2

}

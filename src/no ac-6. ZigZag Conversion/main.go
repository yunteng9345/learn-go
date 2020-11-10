package main

import "fmt"

func main() {
	s := "ABCDEF"
	numRows := 3
	convert(s, numRows)
}

func convert(s string, numRows int) string {
	// 如果是1行返回原字符
	if numRows == 1 {
		return s
	}
	var str string
	total := len(s)
	fmt.Println(total)
	for i := range s {
		fmt.Println(string(s[i]))
	}
	return str
}

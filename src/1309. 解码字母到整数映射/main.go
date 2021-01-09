package main

import (
	"fmt"
	"strings"
)

//给你一个字符串 s，它由数字（'0' - '9'）和 '#' 组成。我们希望按下述规则将 s 映射为一些小写英文字符：
//
//字符（'a' - 'i'）分别用（'1' - '9'）表示。
//字符（'j' - 'z'）分别用（'10#' - '26#'）表示。 
//返回映射之后形成的新字符串。
//
//题目数据保证映射始终唯一。

func main() {
	a := '1'
	b := 'a'
	fmt.Println(b - a)
	fmt.Println(freqAlphabets("2310#11#12"))
}
func freqAlphabets(s string) int {
	count := strings.Count(s, "#")
	res := make([]string, len(s)-2*count)
	for i := 0; i < len(s)-1; i++ {
		j := 0
		if string(s[i]) != "#" {
			res[j] = string(s[i] + 48)
		} else {
			//res[]
		}
		j++
	}
	return strings.Index(s, "#")
}

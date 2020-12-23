package main

import "fmt"

func main() {
	fmt.Println(toLowerCase("Hello"))
}

func toLowerCase(str string) string {
	a := ""
	for i := range str {
		if str[i] >= 65 && str[i] <= 90 {
			a += string(str[i] + 32)
		} else {
			a += string(str[i])
		}
	}
	return a
}

package main

import "fmt"

func main() {
	fmt.Println(calculate("AB"))
}

func calculate(s string) int {
	x, y := 1, 0
	for _, v := range s {
		switch string(v) {
		case "A":
			x = 2*x + y
		case "B":
			y = 2*y + x
		}
	}
	return x + y
}

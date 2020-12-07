package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
}

func numberOfSteps(num int) int {
	res := 0
	for num != 0 {
		if num != 0 && num%2 == 0 {
			num /= 2
			res++
		} else {
			num--
			res++
		}
	}
	return res
}

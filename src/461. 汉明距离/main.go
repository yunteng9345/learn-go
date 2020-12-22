package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	res := 0
	z := x ^ y
	for z != 0 {
		if z&1 == 1 {
			res++
		}
		z >>= 1
	}
	return res
}

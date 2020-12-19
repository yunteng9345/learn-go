package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(printNumbers(1))
}

func printNumbers(n int) []int {
	var list []int
	res := 1
	for float64(res) <= math.Pow10(n)-1 {
		list = append(list, res)
		res++
	}
	return list
}

func mySqrt(x int) int {
	f := float64(x)
	ff := math.Sqrt(f)
	return int(ff)
}

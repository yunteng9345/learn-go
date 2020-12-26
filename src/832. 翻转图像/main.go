package main

import (
	"fmt"
)

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}

func flipAndInvertImage(A [][]int) [][]int {
	for row := range A {
		var temp []int
		for j := len(A[row]) - 1; j >= 0; j-- {
			temp = append(temp, A[row][j]^1)
		}
		A[row] = temp
	}
	return A
}

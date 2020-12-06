package main

import "fmt"

func main() {

	fmt.Println(diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func diagonalSum(mat [][]int) int {
	result := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if i == j {
				result += mat[i][j]
			} else if i+j == len(mat)-1 {
				result += mat[i][j]
			}
		}
	}
	return result
}

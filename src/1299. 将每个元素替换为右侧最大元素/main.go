package main

import (
	"fmt"
)

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
	// 								[ 18, 6 , 6, 6, 1, -1]
}

func replaceElements(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	dp := make([]int, len(arr)+1)
	dp[len(dp)-1] = -1
	for i := len(arr) - 1; i > 0; i-- {
		dp[i] = max(dp[i+1], arr[i])
	}
	return dp[1:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

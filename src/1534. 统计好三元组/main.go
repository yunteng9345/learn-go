package main

import (
	"fmt"
)

func main() {
	//fmt.Println(countGoodTriplets([]int{9, 8, 4, 3}, 1, 2, 3))

	fmt.Println(abs(1, 2))
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	len := len(arr)
	res := 0
	for i := 0; i < len-3; i++ {
		for j := i + 1; i < len-2; j++ {
			if abs(arr[i], arr[j]) <= a {
				for k := j + 1; k <= len-1; k++ {
					if abs(arr[j], arr[k]) <= b && abs(arr[i], arr[k]) <= c {
						res++
					}
				}
			}
		}
	}
	return res
}

// 取绝对值
func abs(a, b int) int {
	if a > b {
		return b - a
	}
	return a - b
}

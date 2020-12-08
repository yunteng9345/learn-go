package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}

func sumOddLengthSubarrays(arr []int) int {
	l := len(arr)
	res := 0
	for i := 1; i <= l; i += 2 {
		for j := 0; j < l; j++ {
			if i+j > l {
				break
			}
			for _, v := range arr[j : j+i] {
				res += v
			}
		}
	}
	return res
}

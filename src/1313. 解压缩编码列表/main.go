package main

import "fmt"

func decompressRLElist(nums []int) []int {
	var result []int
	for i, v := range nums {
		if i%2 != 0 {
			for j := 0; j < nums[i-1]; j++ {
				result = append(result, v)
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 2, 3}
	fmt.Println(decompressRLElist(nums))
}

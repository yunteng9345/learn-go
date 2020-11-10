package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}

func runningSum1(nums []int) []int {
	result := make([]int, len(nums))
	for i := range nums {
		for j := 0; j < i+1; j++ {
			result[i] += nums[j]
		}
	}
	return result
}

func runningSum(nums []int) []int {
	for i := range nums {
		if i != 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}
	return nums
}

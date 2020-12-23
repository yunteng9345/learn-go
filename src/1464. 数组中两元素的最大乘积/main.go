package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}

func maxProduct(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return (nums[0] - 1) * (nums[1] - 1)
}

package main

import "fmt"

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	var nums []int
	for num != 0 {
		nums = append(nums, num%10)
		num = num / 10
	}
	res := 0
	var isReplace bool
	for i := len(nums) - 1; i >= 0; i-- {
		if !isReplace && nums[i] == 6 {
			nums[i] = 9
			isReplace = true
		}
		res = res*10 + nums[i]
	}
	return res
}

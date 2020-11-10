package main

/*
Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

Return the array in the form [x1,y1,x2,y2,...,xn,yn].



Example 1:

Input: nums = [2,5,1,3,4,7], n = 3
Output: [2,3,5,4,1,7]
Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7].
Example 2:

Input: nums = [1,2,3,4,4,3,2,1], n = 4
Output: [1,4,2,3,3,2,4,1]
Example 3:

Input: nums = [1,1,2,2], n = 2
Output: [1,2,1,2]

*/
func main() {
	nums := []int{1, 1, 2, 2}
	n := 2
	shuffle(nums, n)
}

/*
我好菜
*/
func shuffle1(nums []int, n int) []int {
	result := make([]int, len(nums))
	var queue1 []int
	var queue2 []int
	for i := 0; i < len(nums); i++ {
		if i < len(nums)/2 {
			queue1 = append(queue1, nums[i])
		} else {
			queue2 = append(queue2, nums[i])
		}
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			result[i] = queue1[0]
			queue1 = queue1[1:]
		} else {
			result[i] = queue2[0]
			queue2 = queue2[1:]
		}
	}
	return result
}

/*

i'm so stupid...

*/
func shuffle(nums []int, n int) []int {
	result := make([]int, 0, n*2)
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}
	return result
}

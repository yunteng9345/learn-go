package main

func main() {

}

func findNumbers(nums []int) int {
	result := 0
	for i := range nums {
		temp := 0
		for nums[i] != 0 {
			nums[i] /= 10
			temp++
		}
		if temp%2 == 0 {
			result++
		}
	}
	return result
}

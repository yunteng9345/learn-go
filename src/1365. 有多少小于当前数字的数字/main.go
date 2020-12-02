package main

func main() {

}
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				res[i] = res[i] + 1
			}
		}
	}
	return res
}

package main

func main() {
	coins := []int{2, 3, 10}
	minCount(coins)

}
func minCount(coins []int) int {
	var res int
	for _, v := range coins {
		res += (v + 1) / 2
	}
	return res
}

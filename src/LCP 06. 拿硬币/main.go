package main

import "fmt"

func main() {
	coins := []int{2, 3, 10}
	minCount(coins)

}
func minCount(coins []int) int {
	result := 0
	for _, v := range coins {
		result += 1
		if v%2 == 0 {

		}

		fmt.Println(v)
	}
	return result
}

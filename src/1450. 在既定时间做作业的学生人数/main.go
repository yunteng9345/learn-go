package main

import "fmt"

func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
}
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	res := 0
	for i := 0; i < len(startTime); i++ {
		for startTime[i] <= endTime[i] {
			if startTime[i] == queryTime {
				res++
			}
			startTime[i]++
		}
	}
	return res
}

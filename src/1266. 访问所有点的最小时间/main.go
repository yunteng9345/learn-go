package main

import "fmt"

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
}

func minTimeToVisitAllPoints(points [][]int) int {

	for i := 0; i < len(points); i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(points[i][j])
		}
	}
	return 0
}

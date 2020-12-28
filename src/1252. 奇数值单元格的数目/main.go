package main

import "fmt"

func main() {
	n := 2
	m := 3
	indices := [][]int{{0, 1}, {1, 1}}
	fmt.Println(oddCells(n, m, indices))
}

func oddCells(n int, m int, indices [][]int) int {
	grip := make([][]int, n)
	for i := range grip {
		grip[i] = make([]int, m)
	}
	for {

	}
}

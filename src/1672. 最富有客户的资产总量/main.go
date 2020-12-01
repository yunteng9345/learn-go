package main

import "fmt"

func main() {

	account := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println(maximumWealth(account))
}

func maximumWealth(accounts [][]int) int {
	result := 0
	temp := 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			temp = temp + accounts[i][j]
		}
		if temp > result {
			result = temp
		}
		temp = 0
	}
	return result
}

package main

import "fmt"

//Input: J = "aA", S = "aAAbbbb"
//Output: 3
func main() {
	J := "aA"
	S := "aAAbbbb"
	numJewelsInStones(J, S)
}

func numJewelsInStones(J string, S string) int {
	var num int
	for i, v := range S {
		for j := range J {
			if S[i] == J[j] {
				num++
			}
		}
		fmt.Println(string(v))
	}
	return num
}

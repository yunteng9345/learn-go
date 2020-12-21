package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))
}

func judgeCircle(moves string) bool {
	countU := 0
	countD := 0
	countL := 0
	countR := 0
	for i := range moves {
		if moves[i] == 'U' {
			countU++
		}
		if moves[i] == 'D' {
			countD++
		}
		if moves[i] == 'L' {
			countL++
		}
		if moves[i] == 'R' {
			countR++
		}
	}
	return countU-countD == 0 && countL-countR == 0
}

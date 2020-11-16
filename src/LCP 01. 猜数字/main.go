package main

func main() {
	guess := []int{1, 2, 3}
	answer := []int{1, 2, 3}
	game(guess, answer)
}

func game(guess []int, answer []int) int {
	n := 0
	for i := range guess {
		if guess[i] == answer[i] {
			n++
		}
	}
	return n
}

package main

import "fmt"

func main() {
	s := "IV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var num, lastint, total int
	for i := range s {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastint {
			total -= num
		} else {
			total += num
		}
		lastint = num
	}
	return total
}

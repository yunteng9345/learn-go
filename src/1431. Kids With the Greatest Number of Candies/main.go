package main

/*
Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true]

*/
func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	kidsWithCandies(candies, extraCandies)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := 0
	for i, e := range candies {
		if i == 0 || e > m {
			m = e
		}
	}

	result := make([]bool, len(candies))
	for i := range candies {
		if candies[i]+extraCandies >= m {
			result[i] = true
		}
	}
	return result
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	fmt.Println(numEquivDominoPairs(dominoes))

}

func numEquivDominoPairs(dominoes [][]int) int {
	dm := make(map[string]int)
	var key string
	for _, d := range dominoes {
		key = getDominoKey(d)
		if c, ok := dm[key]; ok {
			dm[key] = c + 1
		} else {
			dm[key] = 1
		}
	}
	cnt := 0
	for _, v := range dm {
		cnt += v * (v - 1) / 2
	}
	return cnt
}

func getDominoKey(d []int) string {
	if d[0] <= d[1] {
		return strconv.Itoa(d[0]) + "," + strconv.Itoa(d[1])
	}
	return strconv.Itoa(d[1]) + "," + strconv.Itoa(d[0])
}

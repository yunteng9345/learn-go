package main

import "fmt"

func main() {
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
}

func destCity(paths [][]string) string {
	if len(paths) == 1 {
		return paths[0][1]
	}
	for i := 0; i < len(paths); i++ {
		count := 0
		for j := 0; j < len(paths); j++ {
			if paths[i][1] == paths[j][0] {
				break
			} else {
				count++
			}
			if paths[i][1] != paths[j][0] && count == len(paths) {
				return paths[i][1]
			}
		}
	}
	return ""
}

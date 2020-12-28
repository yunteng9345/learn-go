package main

import "fmt"

func main() {
	//fmt.Println(candy([]int{1, 0, 2}))
	//fmt.Println(candy([]int{1, 2, 2}))
	//fmt.Println(candy([]int{1, 3, 2, 2, 1}))
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))
}

// 每个孩子至少分配到 1 个糖果。
// 相邻的孩子中，评分高的孩子必须获得更多的糖果。
func candy(ratings []int) int {
	res := len(ratings)
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			res++
		}
	}
	return res
}

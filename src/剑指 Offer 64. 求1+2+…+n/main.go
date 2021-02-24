package main

import "fmt"

func main() {
	fmt.Println(sumNums(3))
}
// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

func sumNums(n int) int {
	if n == 1  {
		return 1
	}
	n =  n + sumNums(n-1)
	return n
}
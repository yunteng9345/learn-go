package main

func main() {

}

func transpose(matrix [][]int) [][]int {
	n, m := len(matrix),len(matrix[0])
	t := make([][]int, m)
	for i := range t {
		t[i] = make([]int, n)
		for j := range t[i]{
			t[i][j] = -1
		}
	}
	return matrix
}

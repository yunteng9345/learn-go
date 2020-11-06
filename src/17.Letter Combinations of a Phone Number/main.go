package main

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	//primes := []string{"2", "3", "5", "7", "11", "13"}
	//
	//var s []string = primes[1:4]
	//fmt.Println(s)
	letterCombinations("23")
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	combs := []string{""}
	// 23
	for i := 0; i < len(digits); i++ {
		combsLen := len(combs)
		for j := 0; j < combsLen; j++ {
			for _, letter := range m[digits[i] - '2']{
				combs = append(combs, combs[j] + string(letter))
			}
		}
		combs = combs[combsLen:]
	}
	return combs
}

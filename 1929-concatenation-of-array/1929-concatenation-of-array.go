func getConcatenation(nums []int) []int {
	ans := []int{}
	j := 0
	for j < 2 {
		for _, v := range nums {
			ans = append(ans, v)
		}
		j++
	}

	return ans
}
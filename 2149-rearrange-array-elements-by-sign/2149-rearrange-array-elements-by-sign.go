func rearrangeArray(arr []int) []int {
	ans := make([]int, len(arr))

	pos := 0
	neg := 1

	for _, v := range arr {
		if v >= 0 {
			ans[pos] = v
			pos += 2
		} else {
			ans[neg] = v
			neg += 2
		}
	}

	return ans

}
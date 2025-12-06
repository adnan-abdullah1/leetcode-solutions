func shuffle(nums []int, n int) []int {
	ans := []int{}
	j := n
	for _, v := range nums {
		if j >= len(nums) {
			return ans
		}
		ans = append(ans, v)
		ans = append(ans, nums[j])
		j++
	}

	return ans
}
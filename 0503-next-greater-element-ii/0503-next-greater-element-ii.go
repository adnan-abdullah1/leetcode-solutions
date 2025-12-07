// using virtaul duplicate array (striver)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nge := make([]int, n)

	for i, _ := range nge {
		nge[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n+i-1; j++ { 

			ind := j % n
			if nums[i] < nums[ind] {
				nge[i] = nums[ind]
				break
			}
		}
	}

	return nge
}
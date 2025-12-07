// using virtaul duplicate array (striver)
func BruteNextGreaterElements(nums []int) []int {
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

// optimal with stack using virtaul array
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := []int{}
	ans := make([]int, len(nums))

	popTillNge := func(i int) {
		for len(stack) > 0 && nums[i%n] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
	}

	for i := 2*n - 1; i >= 0; i-- {
		if i >= n {
			popTillNge(i)
		} else {
			popTillNge(i)

			if len(stack) == 0 {
				ans[i] = -1
			} else {
				ans[i] = stack[len(stack)-1]
			}
		}

		stack = append(stack, nums[i%n])

	}

	return ans
}

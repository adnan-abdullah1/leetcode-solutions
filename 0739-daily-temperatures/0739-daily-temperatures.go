func dailyTemperatures(arr []int) []int {
	ans := make([]int, len(arr))
	stack := []int{}

	isEmpty := func() bool {
		return len(stack) == 0
	}
	peek := func() int {
		return stack[len(stack)-1]
	}

	pop := func() {
		stack = stack[:len(stack)-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		for !isEmpty() && arr[i] >= arr[peek()] {
			pop()
		}

		// if stack has still elements, that means warmer day exists
		if !isEmpty() {
			ans[i] = peek() - i
		}

		stack = append(stack, i)
	}

    return ans
}
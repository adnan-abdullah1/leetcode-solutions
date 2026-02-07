func asteroidCollision(arr []int) []int {

	stack := []int{}
	for _, v := range arr {
		if v > 0 {
			stack = append(stack, v)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < int(math.Abs(float64(v))) {
				stack = stack[0 : len(stack)-1]
			}

			if len(stack) > 0 && stack[len(stack)-1] == int(math.Abs(float64(v))) {
				stack = stack[0 : len(stack)-1]
			} else if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, v)
			}
		}
	}

	return stack
}
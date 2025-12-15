func asteroidCollision(arr []int) []int {
	stack := []int{}

	top := func() int {
		return stack[len(stack)-1]
	}
	pop := func() {
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	for i := 0; i < len(arr); i++ {
		isCurrAlive := true
		curr := arr[i]

		for len(stack) > 0 && curr < 0 && top() > 0 {
			t := math.Abs(float64(top()))
			c := math.Abs(float64(curr))

			if t == c {
				pop()
				isCurrAlive = false
                break
			} else if t < c {
				pop()
				isCurrAlive = true
			} else {
				isCurrAlive = false
                break
			}
		}
		if isCurrAlive {
			stack = append(stack, curr)
		}
	}
	return stack
}
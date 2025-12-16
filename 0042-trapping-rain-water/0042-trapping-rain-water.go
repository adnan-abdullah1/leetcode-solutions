func trap(height []int) int {
	stack := []int{}
	water := 0

	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			valleyInd := stack[len(stack)-1] // right side
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			w := i - stack[len(stack)-1] - 1

            water += w * (min(height[i], height[stack[len(stack)-1]]) - height[valleyInd])

		}
        stack=append(stack,i)
	}

	return water
}
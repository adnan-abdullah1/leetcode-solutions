func buildArray(target []int, n int) []string {
	stack := []string{}
	j := 0

	for i := 1; i <= n; i++ {
		stack = append(stack, "Push")

		if target[j] == i {
			j++
			if len(target) == j {
				break
			}
		} else {
			stack = append(stack, "Pop")
		}
	}

	return stack
}
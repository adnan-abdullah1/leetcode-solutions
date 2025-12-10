func removeKdigits(num string, k int) string {
	stack := []rune{}

	for _, v := range num {
		for k > 0 && len(stack) > 0 && v < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, v)
	}

	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	i := 0
	for i < len(stack) && stack[i] == '0' {
		i++
	}

	if i == len(stack) {
		return "0"
	}

	return string(stack[i:])
}
func isValid(s string) bool {
	stack := []string{}

	for _, p := range s {
		if p == '(' || p == '[' || p == '{' {
			stack = append(stack, string(p))
		} else {
			if len(stack) == 0 {
				return false
			}
			el := string(stack[len(stack)-1])
			stack = stack[0 : len(stack)-1]

			if (el == "(" && string(p) != ")") ||
				(el == "[" && string(p) != "]") ||
				(el == "{" && string(p) != "}") {
				return false
			}

		}
	}

	return len(stack) == 0
}
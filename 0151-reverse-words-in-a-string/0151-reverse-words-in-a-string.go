func reverseWords(str string) string {
	ans := ""
	j := len(str) - 1

	for j >= 0 {
		for j >= 0 && str[j] == ' ' {
			j--
		}

		if j < 0 {
			break
		}

		i := j
		for i >= 0 && str[i] != ' ' {
			i--
		}

		subStr := str[i+1 : j+1]
		if len(ans) == 0 {
			ans = ans + subStr
		} else {
			ans = ans + " " + subStr
		}

		j = i
	}
	
    return ans
}
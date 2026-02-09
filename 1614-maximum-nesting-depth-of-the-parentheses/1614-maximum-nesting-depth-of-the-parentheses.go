func maxDepth(str string) int {

	arr := []byte{}

	maxDepth := 0
	for _, v := range str {
		if v == '(' {
			arr = append(arr, '(')
		} else if v == ')' {
			if len(arr) > maxDepth {
				maxDepth = len(arr)
			}

			if len(arr) > 0 {
				arr = arr[0 : len(arr)-1]
			}

		}
	}

	return maxDepth

}
func findMaxConsecutiveOnes(arr []int) int {
	cnt := 0
	maxLen := 0

	for j := range arr {
		if arr[j] == 1 {
			cnt++
		} else {
			cnt = 0
		}
        
		if cnt > maxLen {
			maxLen = cnt
		}
	}

	return maxLen
}
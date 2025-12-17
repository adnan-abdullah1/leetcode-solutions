// tc: n square suffers tle
func BrutelongestOnes(arr []int, k int) int {
	maxi := 0

	for i, _ := range arr {
		zeros := 0

		for j := i; j < len(arr); j++ {
			if arr[j] == 0 {
				zeros++
			}

			if zeros <= k {
				len := j - i + 1
				maxi = max(len, maxi)
			} else {
				break
			}
		}

	}
	return maxi

}
// sliding window o(n)
func longestOnes(arr []int, k int) int {
	maxi := 0
	zeros := 0

	i := 0

	for j := i; j < len(arr); j++ {
		if arr[j] == 0 {
			zeros++
		}

		if zeros <= k {
			len := j - i + 1
			maxi = max(len, maxi)
		} else {
			for zeros > k {
				if arr[i] == 0 {
					zeros--
				}
				i++
			}
		}
	}

	return maxi

}
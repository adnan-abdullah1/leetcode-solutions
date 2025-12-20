func maxScore(arr []int, k int) int {

	lSum := 0
	rSum := 0
	maxi := 0

	for i := 0; i < k; i++ {
		lSum += arr[i]
	}

    maxi = lSum

	j := len(arr) - 1
	for i := k - 1; i >= 0; i-- {
		lSum -= arr[i] 
		rSum += arr[j]
		j--
		maxi = max(maxi, lSum+rSum)
	}

    return maxi
}
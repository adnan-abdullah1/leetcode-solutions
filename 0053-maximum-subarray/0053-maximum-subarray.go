// TLE
func BrutemaxSubArray(arr []int) int {
	maxSum := math.MinInt32
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func maxSubArray(arr []int) int {
	

	sum := arr[0]
    maxSum:=arr[0]

	for j := 1; j < len(arr); j++ {
		
		if sum+arr[j] > arr[j] {
			sum = sum + arr[j]
		} else{
            sum = arr[j]
        }

		if sum > maxSum {
			maxSum = sum
		}
	}
    return maxSum
}
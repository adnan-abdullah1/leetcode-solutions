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

//fails for neg
func FailmaxSubArray(arr []int) int {
	
	sum := 0
    maxSum:=math.MinInt32
	for j := 0; j < len(arr); j++ {
		sum += arr[j]
		if sum < 0 {
			sum = 0
		}

		if sum > maxSum {
			maxSum = sum
		}
	}
    return maxSum
}


func maxSubArray(arr []int) int {

	sum := arr[0]      // best subarray ending at index 0
	maxSum := arr[0]  // best answer so far

	for j := 1; j < len(arr); j++ {
		// this is the line you asked about
		if sum+arr[j] > arr[j] {
			sum = sum + arr[j]
		} else {
			sum = arr[j]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
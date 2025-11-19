func trap(arr []int) int {
    trapped := 0
	for i := range arr {
		rMax := getRightMax(arr, i)
		lMax := getLeftMax(arr, i)
		trapped += min(rMax, lMax) - arr[i]
	}
	return trapped
}


func getLeftMax(arr []int, i int) int {
	maxNo := 0
	for i := i; i >= 0; i-- {
		if arr[i] > maxNo {
			maxNo = arr[i]
		}
	}
	return maxNo
}

func getRightMax(arr []int, i int) int {
	maxNo := 0
	for i := i; i < len(arr); i++ {
		if arr[i] > maxNo {
			maxNo = arr[i]
		}
	}
	return maxNo
}

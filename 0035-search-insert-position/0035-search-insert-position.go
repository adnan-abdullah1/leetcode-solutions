func searchInsert(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	for i := range arr {
		if target <= arr[i] { //1,3,5,6
			return i
		}
	}

	if target > arr[len(arr)-1] {
		return len(arr)
	}

	return -1
}
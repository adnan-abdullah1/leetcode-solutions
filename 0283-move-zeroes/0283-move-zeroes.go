func moveZeroes(arr []int) {
	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

}
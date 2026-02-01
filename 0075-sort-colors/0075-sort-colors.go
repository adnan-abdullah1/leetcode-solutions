func sortColors(arr []int) {
	var l, m int
	var h int = len(arr) - 1

	for m <= h {
		if arr[m] == 0 {
			arr[m], arr[l] = arr[l], arr[m]
			l++
			m++
		} else if arr[m] == 1 {
			m++
		} else {
			arr[m], arr[h] = arr[h], arr[m]
			h--
		}
	}

}
func search(arr []int, key int) int {
	l := 0
	h := len(arr) - 1

	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == key {
			return mid
		}

		// find which part is sorted
		if arr[l] <= arr[mid] {
			// check if key lies in this part
			if key >= arr[l] && key <= arr[mid] {
				h = mid - 1
			} else {
				// key is in right half
				l = mid + 1
			}
		} else {
			if key >= arr[mid] && key <= arr[h] {
				l = mid + 1
			} else {
				// key is in lef half
				h = mid - 1
			}
		}

	}
    return -1
}
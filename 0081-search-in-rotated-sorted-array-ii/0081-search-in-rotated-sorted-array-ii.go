func search(arr []int, key int) bool {
	l := 0
	h := len(arr) - 1

	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == key {
			return true
		}
        // for check when l=m=h
        // [1,0,1,1,1] for key=0
		if arr[l] == arr[mid] && arr[mid] == arr[h] {
			l++
			h--
            continue
		}
		if l > h {
			return false
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
	return false
}
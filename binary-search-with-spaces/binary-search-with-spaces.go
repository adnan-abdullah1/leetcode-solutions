package main

// question on cracking coding interview book pg 107
// binarySearchWithEmpty handles binary search in a sorted string array
// that may contain interspersed empty strings.
func BinarySearchWithEmpty(arr []string, key string) int {
	l, h := 0, len(arr)-1

	for l <= h {
		m := (l + h) / 2

		// If mid is empty, find nearest non-empty string
		if arr[m] == "" {
			L, R := m-1, m+1
			for L >= l || R <= h {
				if L >= l && arr[L] != "" {
					m = L
					break
				}
				if R <= h && arr[R] != "" {
					m = R
					break
				}
				L--
				R++
			}

			// If both sides are exhausted, key not found
			if arr[m] == "" {
				return -1
			}
		}

		if arr[m] == key {
			return m
		}
		if key > arr[m] {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return -1
}

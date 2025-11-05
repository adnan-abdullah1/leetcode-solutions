func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix) - 1 // 0
	cols := len(matrix[0]) - 1 // 1 

	lr := 0 
	hr := rows // 0

	for lr <= hr {
		midRow := (lr + hr) / 2 //0
		if target < matrix[midRow][0] {
			hr = midRow - 1
			continue
		} else if target > matrix[midRow][len(matrix[0])-1] {
			lr = midRow + 1
			continue
		}

		lc := 0
		hc := cols //1

		for lc <= hc {
			midCol := (lc + hc) / 2 // 0, 1

			if matrix[midRow][midCol] == target {
				return true
			} else if target > matrix[midRow][midCol] {
				lc = midCol + 1 // 1
			} else {
				hc = midCol - 1
			}
		}

		return false

	}
	return false
}
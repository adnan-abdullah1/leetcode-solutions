// stair case algo
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	r := 0
	c := len(matrix[0]) - 1

	for r < len(matrix) && c >= 0 {
		if target == matrix[r][c] {
			return true
		}
		// At each step we eliminate exactly one row or one column.
		// Why only one?
		// Because the comparison with matrix[r][c] gives us only one guaranteed fact:
		//
		// - If target > matrix[r][c], all values to the left are smaller,
		//   so the entire current row can be discarded → move down (r++).
		//
		// - If target < matrix[r][c], all values below are larger,
		//   so the entire current column can be discarded → move left (c--).
		//
		// We cannot eliminate both a row and a column at the same time because
		// the comparison only tells us which *one* side is impossible,
		// not both. Eliminating both would remove valid search space.

		if target > matrix[r][c] {
			r++
		} else {
			c--
		}
	}
	return false
}
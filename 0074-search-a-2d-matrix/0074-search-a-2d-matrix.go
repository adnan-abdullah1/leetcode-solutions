// I did this myself gpt cleaned this code

func searchMatrix(matrix [][]int, target int) bool {
	//  Guard clause for empty matrix
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	lowRow, highRow := 0, rows-1

	// \U0001f539 Step 1: Binary search on rows
	for lowRow <= highRow {
		midRow := (lowRow + highRow) / 2

		// If target is smaller than first element in this row → go up
		if target < matrix[midRow][0] {
			highRow = midRow - 1
			continue
		}

		// If target is greater than last element in this row → go down
		if target > matrix[midRow][cols-1] {
			lowRow = midRow + 1
			continue
		}

		// \U0001f539 Step 2: Binary search inside the identified row
		lowCol, highCol := 0, cols-1
		for lowCol <= highCol {
			midCol := (lowCol + highCol) / 2
			if matrix[midRow][midCol] == target {
				return true
			}
			if matrix[midRow][midCol] < target {
				lowCol = midCol + 1
			} else {
				highCol = midCol - 1
			}
		}

		
		return false
	}

	return false
}

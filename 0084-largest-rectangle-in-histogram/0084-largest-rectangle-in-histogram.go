// suffers tle
func BrutelargestRectangleArea(heights []int) int {
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		minH := heights[i]

		for j := i; j < len(heights); j++ {
			if heights[j] < minH {
				minH = heights[j]
			}

			area := minH * (j - i + 1)
			if area > maxArea {
				maxArea = area
			}
		}

	}
	return maxArea
}

// with next smaller left and next smaller right in single pass
func largestRectangleArea(heights []int) int {
    st := []int{} 
    maxArea := 0
    n := len(heights)

    for i := 0; i < n; i++ {
        for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
            tp := st[len(st)-1]   
            st = st[:len(st)-1]   

            var width int
            if len(st) == 0 {
                width = i        
            } else {
                width = i - st[len(st)-1] - 1
            }
            area := width * heights[tp]
            area = max(area,maxArea)
            if area > maxArea {
                maxArea = area
            }
        }
        // push current indexx
        st = append(st, i)
    }

    for len(st) > 0 {
        tp := st[len(st)-1]
        st = st[:len(st)-1]

        var width int
        if len(st) == 0 {
            width = n
        } else {
            width = n - st[len(st)-1] - 1
        }
        area := width * heights[tp]
        if area > maxArea {
            maxArea = area
        }
    }

    return maxArea
}

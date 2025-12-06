// // with help of 2n memroy
// func carFleet(target int, positions []int, speed []int) int {
// 	helper := [][]float64{}

// 	for i, v := range positions {
// 		t := float64(target-v) / float64(speed[i])
// 		helper = append(helper, []float64{float64(v), t})
// 	}

// 	sort.Slice(helper, func(i, j int) bool {
// 		return helper[i][0] > helper[j][0]
// 	})

// 	maxD := 0.0
// 	fleet := 0

// 	for i := 0; i < len(helper); i++ {
// 		if helper[i][1] > maxD {
// 			fleet++
// 			maxD = helper[i][1]
// 		}
// 	}

// 	return fleet
// }

func carFleet(target int, positions []int, speed []int) int {
    n := len(positions)
    
    // Sort positions (and corresponding speeds) descending
    indices := make([]int, n)
    for i := range indices {
        indices[i] = i
    }
    sort.Slice(indices, func(i, j int) bool {
        return positions[indices[i]] > positions[indices[j]]
    })
    
    stack := []float64{}
    
    for _, idx := range indices {
        time := float64(target-positions[idx]) / float64(speed[idx])
        
        if len(stack) == 0 || time > stack[len(stack)-1] {
            stack = append(stack, time)
        }
    }
    
    return len(stack)
}
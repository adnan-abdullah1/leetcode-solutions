func carFleet(target int, positions []int, speed []int) int {
	helper := [][]float64{}

	for i, v := range positions {
		d := float64(target-v) / float64(speed[i])
		helper = append(helper, []float64{float64(v), d})
	}

	sort.Slice(helper, func(i, j int) bool {
		return helper[i][0] > helper[j][0]
	})

	maxD := 0.0
	fleet := 0

	for i := 0; i < len(helper); i++ {
		if helper[i][1] > maxD {
			fleet++
			maxD = helper[i][1]
		}
	}

	return fleet
}

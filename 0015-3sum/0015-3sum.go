func threeSum(nums []int) [][]int {
	ans := [][]int{}
	ansMP := make(map[string]int, len(nums))

	for i := 0; i < len(nums)-2; i++ {
		// fix for [0,0,0,0............]
		// otherwise without this tle
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		mp := make(map[int]int, len(nums))

		for j := i + 1; j < len(nums); j++ {
			need := 0 - nums[i] - nums[j]

			ind, ok := mp[need]
			if ok {
				triplet := []int{nums[i], nums[j], nums[ind]}

				sort.Ints(triplet)

				hashKey := fmt.Sprintf("%d-%d-%d", triplet[0], triplet[1], triplet[2])
				_, ok := ansMP[hashKey]
				if !ok {
					ansMP[hashKey] = 1
					ans = append(ans, triplet)
				}

			} else {
				mp[nums[j]] = j
			}
		}
	}

	return ans
}

// func threeSum(nums []int) [][]int {
// 	ans := [][]int{}
// 	mp := make(map[string]bool, len(nums))

// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			for l := j + 1; l < len(nums); l++ {
// 				if nums[i]+nums[j]+nums[l] == 0 {

// 					m := []int{nums[i], nums[j], nums[l]}
// 					sort.Ints(m) // -1,0,1

// 					str := fmt.Sprintf("%d %d %d", m[0], m[1], m[2])

// 					_, ok := mp[str]
// 					if !ok {
// 						mp[str] = true
// 						ans = append(ans, []int{m[0], m[1], m[2]})
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return ans
// }

// better with set for duplicate check
func threeSum(nums []int) [][]int {
    results := [][]int{}
    seenTriplets := make(map[string]bool) // stores unique triplet signatures

    for i := 0; i < len(nums)-2; i++ {

        // skip duplicate
        // ex: [0,0,0,0,0,0,0]
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        valueToIndex := make(map[int]int) // used for two-sum lookup

        for j := i + 1; j < len(nums); j++ {
            needed := 0 -(nums[i] + nums[j]) // value needed to complete triplet

            if kIndex, found := valueToIndex[needed]; found {

                triplet := []int{nums[i], nums[j], nums[kIndex]}
                sort.Ints(triplet)

                key := fmt.Sprintf("%d-%d-%d", triplet[0], triplet[1], triplet[2])

                if !seenTriplets[key] {
                    seenTriplets[key] = true
                    results = append(results, triplet)
                }

            } else {
                valueToIndex[nums[j]] = j
            }
        }
    }

    return results
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
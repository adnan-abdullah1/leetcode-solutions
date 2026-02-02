
func BrutemajorityElement(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}

	for k, v := range mp {
		if v > len(nums)/22 {
			return k
		}
	}

	return -1
}

// buyers moore voting alog
func majorityElement(nums []int) int {
	cand := 0
	count := 0

	for _, v := range nums {
		if count == 0 {
			cand = v
		}

		if cand != v {
			count--
		} else {
			count++
		}
	}

	return cand
}
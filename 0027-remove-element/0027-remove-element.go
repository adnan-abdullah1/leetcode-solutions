
// with extra mem
// func removeElement(nums []int, val int) int {
// 	ans := []int{}

// 	for _, v := range nums {

// 		if v != val {
// 			ans = append(ans, v)
// 		}
// 	}

// 	copy(nums, ans)
// 	return len(ans)

// }

func removeElement(nums []int, val int) int {
	k := 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}

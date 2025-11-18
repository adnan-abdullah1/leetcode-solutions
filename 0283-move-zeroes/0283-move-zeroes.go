func moveZeroes(nums []int) {
	for i, v := range nums {

		if v != 0 {
			continue
		}

		j := i
		for j<len(nums) && nums[j] == 0 {
			j++
		}

		if j >= len(nums) {
			break
		}

        nums[i],nums[j]=nums[j],nums[i]

	}
} 
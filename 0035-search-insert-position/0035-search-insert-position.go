func searchInsert(nums []int, target int) int {
	var l, h = 0, len(nums) - 1

	for l <= h {
		mid := (l + h) / 2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			h = mid - 1

		}
	}
	return l
}
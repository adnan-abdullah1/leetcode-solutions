// both have same complexity

// func moveZeroes(nums []int) {
//     n := len(nums)

//     for i := 0; i < n; i++ {

//         // if it's non-zero, just go ahead
//         if nums[i] != 0 {
//             continue
//         }

//         // find next non-zero
//         j := i + 1
//         for j < n && nums[j] == 0 {
//             j++
//         }

//         // no more non-zeros
//         if j == n {
//             return
//         }

//         // swap next non-zero into current position
//         nums[i], nums[j] = nums[j], nums[i]
//     }
// }

func moveZeroes(nums []int) {
	n := len(nums)

	pos := 0
	for i := 0; i < n; i++ {

		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}
	// now fill zeros at rest end
	for pos < len(nums) {
		nums[pos] = 0
		pos++
	}
}

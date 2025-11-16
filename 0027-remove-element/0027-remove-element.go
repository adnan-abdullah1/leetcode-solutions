// func removeElement(nums []int, val int) int {
// 	// [5,6,3,3,,3,1,3,8,7]
//     k:=1
//     j := 0
// 	for i := 0; i < len(nums); i++ {
//         if nums[i] != val{
//             continue
//         }

//         for j<len(nums) && val == nums[j]{
//             j++
//         }

//         if j >= len(nums){
//             break;
//         }

//         nums[i]=nums[j];
//         k++
// 	}
//     return k
// }

// with extra mem
func removeElement(nums []int, val int) int {
	ans := []int{}

	for _, v := range nums {

		if v != val {
			ans = append(ans, v)
		}
	}

	copy(nums, ans)
	return len(ans)

}
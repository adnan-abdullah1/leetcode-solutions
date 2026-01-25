func subsets(nums []int) [][]int {
	res := [][]int{}
	ds := []int{}

	return sets(nums, 0, ds, res)
}

func sets(nums []int, i int, ds []int, res [][]int) [][]int {
	if i == len(nums) {
		tmp := make([]int, len(ds))
		copy(tmp, ds)
		res = append(res, tmp)
		return res
	}

	ds = append(ds, nums[i])

	res=sets(nums, i+1, ds, res)

	ds = ds[:len(ds)-1]
	res=sets(nums, i+1, ds, res)
	return res
}

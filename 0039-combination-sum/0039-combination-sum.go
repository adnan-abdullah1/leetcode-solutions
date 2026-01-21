func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	findCombs(0, target, candidates, &ans, []int{})
	return ans
}

func findCombs(ind, target int, arr []int, ans *[][]int, ds []int) {
	if ind == len(arr) {
		if target == 0 {
			temp := make([]int, len(ds))
			copy(temp, ds)
			(*ans) = append(*ans, temp)
		}
		return
	}

	if arr[ind] <= target {
		ds = append(ds, arr[ind])
		findCombs(ind, target-arr[ind], arr, ans, ds)
		ds = ds[:len(ds)-1]
	}
	// skip
	findCombs(ind+1, target, arr, ans, ds)
}
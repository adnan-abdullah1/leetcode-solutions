func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for _, num := range nums { // num is 1,2
		size := len(res)      // 1,2
		for i := range size { // i=0, i=0
			tmp := append([]int{}, res[i]...) // [1,2]
			tmp = append(tmp, num)            //[1]
			res = append(res, tmp)            // [{},[1]]
		}
	}
	return res
}

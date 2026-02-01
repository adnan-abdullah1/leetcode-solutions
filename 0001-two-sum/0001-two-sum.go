func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		find := target - v
		n, ok := mp[find]
		if ok {
			return []int{n, i}
		} else {
			mp[v] = i
		}
	}
	return []int{}
}
func subarraySum(arr []int, k int) int {
	mp := make(map[int]int)
    mp[0]=1
	s := 0
	total := 0

	for i := 0; i < len(arr); i++ {
		s += arr[i]
		
			// does s-k exist in map
			f, v := mp[s-k]
			if v {
				total+=f
			}
		mp[s]++

	}
	return total
}
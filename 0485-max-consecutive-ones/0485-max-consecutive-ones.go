func findMaxConsecutiveOnes(nums []int) int {
	cnt := 0
	maxCnt := 0

	for _, v := range nums {
		if v != 0 {
			cnt++
			if cnt > maxCnt {
				maxCnt = cnt
			}
		} else {
			cnt = 0
		}
	}
	return maxCnt
}
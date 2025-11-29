func fourSum(arr []int, target int) [][]int {
	sort.Ints(arr)
	ans := [][]int{}

	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < len(arr); j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			k := j + 1
			l := len(arr) - 1

			for k < l {
				total := arr[i] + arr[j] + arr[k] + arr[l]
				if total < target {
					k++
				} else if total > target {
					l--
				} else {
					ans = append(ans, []int{arr[i], arr[j], arr[k], arr[l]})
					k++
					l--
					for k < l && arr[k] == arr[k-1] {
						k++
					}
					for k < l && arr[l] == arr[l+1] {
						l--
					}
				}
			}
		}
	}
	return ans
}
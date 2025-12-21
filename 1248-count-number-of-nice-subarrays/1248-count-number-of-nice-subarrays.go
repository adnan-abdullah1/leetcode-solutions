// suffers tle
func BrutenumberOfSubarrays(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		oddCnt := 0
		for j := i; j < len(arr); j++ {
			if arr[j]%2 != 0 {
				oddCnt++
			}

			if oddCnt > k {
				break
			}

			if oddCnt == k {
				count++
			}
		}
	}
	return count
}

func numberOfSubarrays(arr []int, k int) int {
    return atMostK(arr, k) - atMostK(arr, k-1)
}


func atMostK(arr []int, k int) int {
    i := 0
    count := 0
    oddCnt := 0

    for j := 0; j < len(arr); j++ {
        if arr[j]%2 != 0 {
            oddCnt++
        }

        for oddCnt > k {
            if arr[i]%2 != 0 {
                oddCnt--
            }
            i++
        }

        count += j - i + 1
    }
    return count
}

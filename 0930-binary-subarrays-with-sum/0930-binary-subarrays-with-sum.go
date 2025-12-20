// taking map with prefix sum

func Approach1numSubarraysWithSum(arr []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	s := 0
	total := 0

	for i := 0; i < len(arr); i++ {
		s += arr[i]
		total += mp[s-k]
		mp[s]++
	}
	return total

}

// with set decompostion
//

// taking map with prefix sum
// count of all sub array with sum = goal
// is equal to count of all sub array with sum <= goal - count of all sub aray with sum <= goal-1
// then with this we are able to count all
func numSubarraysWithSum(arr []int, k int) int {
    return atMost(arr, k) - atMost(arr, k-1)
}


func atMost(arr []int, limit int) int {
    if limit < 0 {
        return 0
    }

    left := 0
    sum := 0
    count := 0

    for right := 0; right < len(arr); right++ {
        sum += arr[right]

        for sum > limit {
            sum -= arr[left]
            left++
        }

        count += right - left + 1
    }

    return count
}

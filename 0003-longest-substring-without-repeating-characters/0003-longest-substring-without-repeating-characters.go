func lengthOfLongestSubstring(s string) int {
	b := []rune(s)

	seen := make(map[rune]bool)
	i := 0
	maxLen := 0

	for j := 0; j < len(b); j++ {
		for seen[b[j]] {
			delete(seen, b[i])
			i++
		}
		seen[b[j]] = true
		if j-i+1 > maxLen {
			maxLen = j - i + 1
		}
	}
	return maxLen
}
func isSubsequence(s1 string, s2 string) bool {
	var i int
	for j := 0; j < len(s1); j++ {
		for i < len(s2) && s2[i] != s1[j] {
			i++
		}
		if i >= len(s2) || s2[i] != s1[j] {
			return false
		}
        if s2[i] == s1[j] {
			i++
		}
	}
	return true
}
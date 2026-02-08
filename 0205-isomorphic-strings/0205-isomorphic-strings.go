func isIsomorphic(s, t string) bool {
	mp := make(map[string]string)
	for i, v := range s {
		val, ok := mp[string(v)]
		if ok {
			if string(t[i]) != val {
				return false
			}
		} else {
			mp[string(v)] = string(t[i])

		}
	}
	mp = make(map[string]string)
	for i, v := range t {
		val, ok := mp[string(v)]
		if ok {
			if string(s[i]) != val {
				return false
			}
		} else {
			mp[string(v)] = string(s[i])

		}
	}
	return true
}

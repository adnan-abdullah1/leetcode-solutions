func groupAnagrams(mainStr []string) [][]string {

	ans := [][]string{}
	mp := make(map[string][]string)
	for _, m := range mainStr {
		freq := countFreq(m)

		mp[freq] = append(mp[freq], m)
	}

	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func countFreq(str string) string {
	arr := make([]int, 52)

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			arr[c-'a']++ // 98-97
		} else if c >= 'A' && c <= 'Z' {
			arr[c-'A'+26]++ // 66-65+26
		}
	}

	
	parts := make([]string, len(arr))
	for i, v := range arr {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, "#")
}





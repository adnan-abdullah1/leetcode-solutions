func frequencySort(str string) string {
	mp := make(map[string]int, len(str))

	type Pair struct {
		ch  string
		cnt int
	}

	pairs := make([]Pair, 0, len(mp))

	for _, v := range str {
		val, _ := mp[string(v)]
		mp[string(v)] = val + 1
	}

	for k, v := range mp {
		pairs = append(pairs, Pair{
			ch:  k,
			cnt: v,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].cnt > pairs[j].cnt
	})

    ans := ""
	for _, strct := range pairs {
		ans += strings.Repeat(strct.ch, strct.cnt)
	}

	return ans
}
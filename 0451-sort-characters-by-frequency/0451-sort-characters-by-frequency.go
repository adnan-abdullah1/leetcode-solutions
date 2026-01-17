func frequencySort(s string) string {

	mp := make(map[int]int)

	for _, v := range s {
		mp[int(v)]++
	}

	type Pair struct {
		ch   int
		freq int
	}

	pairs := make([]Pair, 0, len(mp))

	for k, v := range mp {
		pairs = append(pairs, Pair{
			ch:   k,
			freq: v,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	ans := ""

	for _, v := range pairs {
		ans += strings.Repeat(string(v.ch), v.freq)
	}

	return ans
}
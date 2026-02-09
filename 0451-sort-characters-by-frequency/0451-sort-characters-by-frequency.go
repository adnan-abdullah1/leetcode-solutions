func frequencySort(str string) string {
	mp := make(map[string]int)

	for _, v := range str {
		s := string(v)
		val, _ := mp[s]
		mp[s] = val + 1
	}

	type Pair struct {
		ch  string
		cnt int
	}

	paris := make([]Pair, 0, len(mp))

	for k, v := range mp {
		paris = append(paris, Pair{
			ch:  k,
			cnt: v,
		})
	}

	sort.Slice(paris, func(i, j int) bool {
		return paris[i].cnt > paris[j].cnt
	})

	ans := ""
	for _, v := range paris {
		ans += strings.Repeat(v.ch, v.cnt)
	}

	return ans
}
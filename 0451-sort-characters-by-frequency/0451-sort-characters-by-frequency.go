
func BrutefrequencySort(s string) string {

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


type KV struct {
	ch   int
	freq int
}
type MaxHeap []KV

func (h *MaxHeap) Push(x any) {
	*h = append(*h, KV{
		ch:   x.(KV).ch,
		freq: x.(KV).freq,
	})
}

func (h *MaxHeap) Pop() any {
	d := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return d
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i].freq > (*h)[j].freq
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}


func frequencySort(s string) string {
	mp := make(map[int]int)

	for _, v := range s {
		mp[int(v)]++
	}

	type Pair struct {
		ch   int
		freq int
	}

	h := &MaxHeap{}
	heap.Init(h)

	for k, f := range mp {
		heap.Push(h, KV{
			ch:   k,
			freq: f,
		})

	}

	var ans strings.Builder

	for h.Len() > 0 {
		pair := heap.Pop(h)
		ans.WriteString(strings.Repeat(string(rune(pair.(KV).ch)), pair.(KV).freq))
	}

	return ans.String()
}
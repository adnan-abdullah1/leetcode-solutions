type Item struct {
	no   int
	freq int
}
type MaxHeap []Item

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i].freq > (*h)[j].freq
}
func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}


func topKFrequent(arr []int, k int) []int {

	mp := make(map[int]int)
	for _, v := range arr {
		mp[v]++
	}

    h := &MaxHeap{}
	heap.Init(h)
    
	// heapify
	for n, f := range mp {
		heap.Push(h, Item{
			no:   n,
			freq: f,
		})
	}

	ans := []int{}
	for k > 0 && h.Len() > 0 {
		x := heap.Pop(h).(Item)
		ans = append(ans, x.no)
		k--
	}

	return ans
}
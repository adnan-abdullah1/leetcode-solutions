type MaxHeap []int

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	if len(*h) == 0 {
		return -1
	}

	d := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return d
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

// Swap swaps the elements with indexes i and j.
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}

	heap.Init(h)

	for _, v := range stones {
		heap.Push(h, v)
	}

	for h.Len() > 1 {
		v1 := heap.Pop(h).(int) 
		v2 := heap.Pop(h).(int) 

		if v1 != v2 {
			heap.Push(h, v1-v2)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}


type MinHeap []int

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	if len(*h) == 0 {
		return -1
	}
	el := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return el
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func nthUglyNumber(n int) int {
	seen := map[int]bool{1: true}
	arr := []int{1}

	h := MinHeap(arr)
	heap.Init(&h)

	curr := 0
	for h.Len() > 0 && n > 0 {
		curr = heap.Pop(&h).(int)

		v1 := curr * 2
		if !seen[v1] {
			seen[v1] = true
			heap.Push(&h, v1)
		}

		v2 := curr * 3
		if !seen[v2] {
			seen[v2] = true
			heap.Push(&h, v2)
		}

		v3 := curr * 5
		if !seen[v3] {
			seen[v3] = true
			heap.Push(&h, v3)
		}
        
		n--
	}
	return curr
}


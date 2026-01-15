type MaxHeap []int

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	if len(*h) == 0 {
		return -1
	}
	el := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return el
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func findKthLargest(nums []int, k int) int {
	h := MaxHeap(nums)

	heap.Init(&h)

	for h.Len() > 0 {
		d := heap.Pop(&h)
		if k == 1 {
			return d.(int)
		}
		k--
	}

	return -1
}
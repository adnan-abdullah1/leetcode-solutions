type T struct {
	Sum int
	i   int
	j   int
}
type MaxHeap []T

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(T))
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
	return (*h)[i].Sum > (*h)[j].Sum
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Swap(i, j int) {

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Top() any {
	if len(*h) == 0 {
		return -1
	}
	return (*h)[0]
}

func kSmallestPairs(arr1 []int, arr2 []int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)

	pushIntoHeap := func(i, j, sum int) {
		heap.Push(h, T{
			Sum: sum,
			i:   i,
			j:   j,
		})
	}

	for i, v1 := range arr1 {
		for j, v2 := range arr2 {
			if h.Len() < k {
				pushIntoHeap(i, j, v1+v2)

			} else { // heap has more or equal to k
				if h.Top().(T).Sum > v1+v2 {
					heap.Pop(h)
					pushIntoHeap(i, j, v1+v2)
				} else {
					break
				}
			}
		}
	}

	ans := [][]int{}

	for h.Len() > 0 {
		el := heap.Pop(h).(T)
		ans = append(ans, []int{arr1[el.i], arr2[el.j]})
	}

	return ans
}
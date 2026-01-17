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

// n*m*log(k)
func BrutekSmallestPairs(arr1 []int, arr2 []int, k int) [][]int {
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

// --------------------------------------------------------------------------------------------------------------------

type PairSum struct {
	Sum int
	i   int
	j   int
}

type MinHeap []PairSum

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(PairSum))
}

func (h *MinHeap) Pop() any {
	d := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return d
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].Sum < (*h)[j].Sum
}

func (h *MinHeap) Len() int {
	return len(*h)
}

// klog(k)
func kSmallestPairs(arr1, arr2 []int, k int) [][]int {

	h := &MinHeap{}
	heap.Init(h)

	type key struct{ i, j int }
	visited := make(map[struct{ i, j int }]bool)

	ans := [][]int{}

	pushToHeap := func(i, j, sum int) {
		heap.Push(h, PairSum{
			Sum: sum,
			i:   i,
			j:   j,
		})
	}

	if len(arr1) == 0 && len(arr2) == 0 {
		return ans
	}

	heap.Push(h, PairSum{
		Sum: arr1[0] + arr2[0],
		i:   0,
		j:   0,
	})

	for k > 0 && h.Len() > 0 {
        
		d := heap.Pop(h).(PairSum)
		ans = append(ans, []int{arr1[d.i], arr2[d.j]})

		if d.i+1 < len(arr1) {

			i := d.i + 1
			j := d.j

			if !visited[key{i: i, j: j}] {
                pushToHeap(i,j,arr1[i] + arr2[j])
				visited[key{i: i, j: j}] = true
			}
		}

		if d.j+1 < len(arr2) {

			i := d.i
			j := d.j + 1

			if !visited[key{i: i, j: j}] {
				pushToHeap(i,j,arr1[i] + arr2[j])
				visited[key{i: i, j: j}] = true
			}

		}

		k--
	}

	return ans
}
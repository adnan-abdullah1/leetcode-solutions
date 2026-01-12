func lastStoneWeight(stones []int) int {
	inverseIntComparator := func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	}
	heap := binaryheap.NewWith(inverseIntComparator)

	for _, v := range stones {
		heap.Push(v)
	}

	for heap.Size() > 1 {
		n1i, _ := heap.Pop()
		n2i, _ := heap.Pop()

		n1 := n1i.(int)
		n2 := n2i.(int)
		if n1 == n2 {
			continue
		}
		rem := n1 - n2
		if rem < 0 {
			rem = -1 * rem
		}
		heap.Push(rem)
	}

	if heap.Size() == 0 {
		return 0
	}
	maxi, _ := heap.Peek()
	return maxi.(int)
}
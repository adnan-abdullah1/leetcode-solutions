type MinStack struct {
	stack []int
	small []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		small: []int{},
	}
}

func (this *MinStack) Push(val int) {
	// push to stack
	this.stack = append(this.stack, val)
	// when small array is empty then push direct val
	if len(this.small) == 0 {
		this.small = append(this.small, val)
		return
	}
	// otherwise check last min el in small if it is greater than val
	if val <= this.small[len(this.small)-1] {
		this.small = append(this.small, val)
	}
}

func (this *MinStack) Pop() {

	if len(this.stack) > 0 {
		top := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]

		if top == this.small[len(this.small)-1] {
			this.small = this.small[:len(this.small)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.small) > 0 {
		return this.small[len(this.small)-1]
	}

	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
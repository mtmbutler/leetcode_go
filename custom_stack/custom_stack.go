package customstack

// CustomStack is an implementation of the stack paradigm.
type CustomStack struct {
	maxSize int
	stack   []int
}

// Constructor initializes a CustomStack object with maxSize which is the maximum
// number of elements in the stack or do nothing if the stack reached the maxSize.
func Constructor(maxSize int) CustomStack {
	return CustomStack{maxSize: maxSize}
}

// Push adds x to the top of the stack if the stack hasn't reached the maxSize.
func (this *CustomStack) Push(x int) {
	if len(this.stack) < this.maxSize {
		this.stack = append(this.stack, x)
	}
}

// Pop pops and returns the top of stack or -1 if the stack is empty.
func (this *CustomStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	lastIx := len(this.stack) - 1
	result := this.stack[lastIx]
	this.stack = this.stack[:lastIx]
	return result
}

// Increment increments the bottom k elements of the stack by val. If there are
// less than k elements in the stack, just increment all the elements in the
// stack.
func (this *CustomStack) Increment(k int, val int) {
	numElements := k
	if len(this.stack) < k {
		numElements = len(this.stack)
	}
	for i := 0; i < numElements; i++ {
		this.stack[i] += val
	}

}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

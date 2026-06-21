type MinStack struct {
	elements []int
}

func Constructor() MinStack {
	return MinStack{
		elements: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
}

func (this *MinStack) Pop() {
	this.elements = this.elements[:len(this.elements) - 1]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements) - 1]
}

func (this *MinStack) GetMin() int {
    m := this.elements[0]

    for _,n := range this.elements {
        m = min(m,n)
    }

    return m
}
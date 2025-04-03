package funcs

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	list := NewList(head)
	for i := 0; i < k; i++ {
		list.TailToHead()
	}

	return list.Head()
}

type List struct {
	head *ListNode
	size int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(head *ListNode) *List {
	list := &List{head: head}
	list.setSize()

	return list
}

func (this *List) setSize() {
	size := 1
	next := this.head
	for next.Next != nil {
		size++
		next = next.Next
	}
	this.size = size
}

func (this *List) Size() int {
	return this.size
}

func (this *List) Head() *ListNode {
	return this.head
}

func (this *List) Tail() *ListNode {
	next := this.head
	for next.Next != nil {
		next = next.Next
	}

	return next
}

func (this *List) ToSlice() []int {
	s := make([]int, 0, this.Size())
	next := this.head
	for next.Next != nil {
		s = append(s, next.Val)
		next = next.Next
	}
	s = append(s, next.Val)

	return s
}

func (this *List) Push(value int) int {
	tail := this.Tail()
	tail.Next = &ListNode{Val: value, Next: nil}
	this.size++

	return this.Size()
}

func (this *List) Pop() *ListNode {
	if this.Size() == 1 {
		return nil
	}

	preTail := new(ListNode)
	tail := this.Tail()

	next := this.head
	for next.Next != nil {
		if next.Next == tail {
			preTail = next
		}
		next = next.Next
	}
	preTail.Next = nil
	this.size--

	return tail
}

func (this *List) TailToHead() {
	if this.Size() == 1 {
		return
	}
	tail := this.Pop()
	tail.Next = this.head
	this.head = tail
	this.size++
}

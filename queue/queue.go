package queue

// An FIFO queue
type Queue []int

// Push element into queue
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop element from head
// 	e.g. Pop(1)
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
package data

type Queue struct {
	ll *LinkedList
}

func NewQueue() *Queue {
	return &Queue{ll: &LinkedList{}}
}

func (q *Queue) Push(val interface{}) {
	q.ll.AddNode(val)
}

func (q *Queue) Pop() interface{} {
	front := q.ll.Front()
	q.ll.PopFront()
	return front
}

func (q *Queue) Empty() bool {
	return q.ll.Empty()
}
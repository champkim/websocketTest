package types

type Queue struct {
	ll    *LinkedList
	count int
}

func NewQueue() *Queue {
	return &Queue{ll: &LinkedList{}, count: 0}
}

func (q *Queue) Push(val interface{}) {
	q.ll.AddNode(val)
	q.count++
}

func (q *Queue) Pop() interface{} {
	front := q.ll.Front()
	q.ll.PopFront()
	q.count--
	return front
}

func (q *Queue) Empty() bool {
	return q.ll.Empty()
}

func (q *Queue) Count() int {
	return q.count
}

func ChangeQueue(srcQ **Queue, targetQ **Queue) {
	// fmt.Printf("src %v ->", *srcQ)
	// fmt.Printf("target %v ->", *targetQ)

	var tmpQ *Queue = *srcQ
	*srcQ = *targetQ
	*targetQ = tmpQ

	// fmt.Println("")
	// fmt.Printf("changed src %v ->", *srcQ)
	// fmt.Printf("changed target %v ->", *targetQ)

	// var tmpQ *data.Queue = queue1
	// queue1 = queue2
	// queue2 = tmpQ

	// fmt.Printf("changed src %v ->", srcQ)
	// fmt.Printf("changed target %v ->", targetQ)
	// fmt.Println("<<<<<<<<<<<<< srcQ >>>>>>>>>>>>>>")
	// fmt.Printf("srcQ %v ->", srcQ)
	// for !srcQ.Empty() {
	//   val := srcQ.Pop()
	//   fmt.Printf("%s ->", val)
	// }

	// fmt.Println("<<<<<<<<<<<<< targetQ >>>>>>>>>>>>>>")
	// fmt.Printf("targetQ %v ->", targetQ)
	// for !targetQ.Empty() {
	//   val := targetQ.Pop()
	//   fmt.Printf("%s ->", val)
	// }
}
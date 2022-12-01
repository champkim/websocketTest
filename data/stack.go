package data

type Stack struct {
	ll *LinkedList
}

func NewStack() *Stack {
	return &Stack{ll: &LinkedList{}}
}

func (s *Stack) Empty() bool {
	return s.ll.Empty()
}

func (s *Stack) Push(val interface{}) {
	s.ll.AddNode(val)
}

func (s *Stack) Pop() interface{} {
	back := s.ll.Back()
	s.ll.PopBack()
	return back
}
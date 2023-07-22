package collections

type (
	Stack struct {
		*linkedList
	}
)

func NewStack() *Stack {
	return &Stack{
		linkedList: &linkedList{},
	}
}

// Push -
// O (1)
func (s *Stack) Push(value int) {
	s.insertEnd(newLinkedListNode(value))
}

// Pop -
// O (1)
func (s *Stack) Pop() (int, error) {
	if s.tail == nil {
		return 0, ErrEmpty
	}

	value := s.tail.value

	s.deleteEnd()

	return value, nil
}

func (s *Stack) Print() {
	s.printReverse()
}

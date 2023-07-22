package collections

type (
	Queue struct {
		*linkedList
	}
)

func NewQueue() *Queue {
	return &Queue{
		linkedList: &linkedList{},
	}
}

// Add -
// O (1)
func (s *Queue) Add(value int) {
	s.insertEnd(newLinkedListNode(value))
}

// Get -
// O (1)
func (s *Queue) Get() (int, error) {
	if s.head == nil {
		return 0, ErrEmpty
	}

	value := s.head.value

	s.deleteStart()

	return value, nil
}

func (s *Queue) Print() {
	s.print()
}

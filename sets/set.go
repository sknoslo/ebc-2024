package sets

type empty struct{}

var E = empty{}

type Set[T comparable] struct {
	items map[T]empty
}

func New[T comparable](size int) *Set[T] {
	return &Set[T]{
		items: make(map[T]empty, max(0, size)),
	}
}

func (s *Set[T]) Insert(item T) {
	s.items[item] = E
}

func (s *Set[T]) Has(item T) bool {
	_, ok := s.items[item]

	return ok
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

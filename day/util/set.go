package util

type Set[T comparable] struct {
	mapSet map[T]struct{}
	size   int
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		mapSet: make(map[T]struct{}),
		size:   0,
	}
}

func (s *Set[T]) Empty() bool {
	return s.size == 0
}

func (s *Set[T]) Add(item T) {
	if _, ok := s.mapSet[item]; !ok {
		s.mapSet[item] = struct{}{}
		s.size++
	}
}

func (s *Set[T]) Remove(item T) {
	if _, ok := s.mapSet[item]; ok {
		delete(s.mapSet, item)
		s.size--
	}
}

func (s *Set[T]) Contains(item T) bool {
	if _, ok := s.mapSet[item]; ok {
		return true
	}
	return false
}

func (s *Set[T]) Size() int {
	return s.size
}

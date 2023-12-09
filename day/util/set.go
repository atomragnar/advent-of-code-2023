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

func SetFromSlice[T comparable](s []T) *Set[T] {
	set := NewSet[T]()
	for _, item := range s {
		set.Add(item)
	}
	return set
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

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	return nil
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	return nil
}

func Difference[T comparable](first, second *Set[T]) *Set[T] {
	return nil
}

func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	return nil
}

func IsSubset[T comparable](set, subSet *Set[T]) bool {
	for item := range subSet.mapSet {
		if _, ok := set.mapSet[item]; !ok {
			return ok
		}
	}
	return true
}

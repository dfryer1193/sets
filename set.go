package sets

type Set[Type comparable] struct {
	set map[Type]struct{}
}

func New[Type comparable]() *Set[Type] {
	return &Set[Type]{
		set: make(map[Type]struct{}),
	}
}

func From[Type comparable](values ...Type) *Set[Type] {
	set := New[Type]()
	for _, value := range values {
		set.Add(value)
	}
	return set
}

func (s *Set[Type]) Add(value Type) {
	s.set[value] = struct{}{}
}

func (s *Set[Type]) Has(value Type) bool {
	_, ok := s.set[value]
	return ok
}

func (s *Set[Type]) Remove(value Type) {
	delete(s.set, value)
}

func (s *Set[Type]) Size() int {
	return len(s.set)
}

func (s *Set[Type]) Union(other *Set[Type]) *Set[Type] {
	union := New[Type]()
	for value := range s.set {
		union.Add(value)
	}
	for value := range other.set {
		union.Add(value)
	}
	return union
}

func (s *Set[Type]) Intersection(other *Set[Type]) *Set[Type] {
	intersection := New[Type]()
	for value := range s.set {
		if other.Has(value) {
			intersection.Add(value)
		}
	}
	return intersection
}

func (s *Set[Type]) Difference(other *Set[Type]) *Set[Type] {
	difference := New[Type]()
	for value := range s.set {
		if !other.Has(value) {
			difference.Add(value)
		}
	}
	return difference
}

func (s *Set[Type]) SymmetricDifference(other *Set[Type]) *Set[Type] {
	difference := New[Type]()
	for value := range s.set {
		if !other.Has(value) {
			difference.Add(value)
		}
	}
	for value := range other.set {
		if !s.Has(value) {
			difference.Add(value)
		}
	}
	return difference
}

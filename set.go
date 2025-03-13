package sets

// Set is a generic type representing a collection of unique elements of a comparable type.
type Set[Type comparable] struct {
	set map[Type]struct{}
}

// New initializes and returns a pointer to an empty Set of a given comparable type.
func New[Type comparable]() *Set[Type] {
	return &Set[Type]{
		set: make(map[Type]struct{}),
	}
}

// From creates and returns a new Set initialized with the provided values of a comparable type.
func From[Type comparable](values ...Type) *Set[Type] {
	set := New[Type]()
	for _, value := range values {
		set.Add(value)
	}
	return set
}

// Add inserts a unique element into the set; it has no effect if the element already exists.
func (s *Set[Type]) Add(value Type) {
	s.set[value] = struct{}{}
}

// Has checks if the specified value exists in the set, returning true if found or false otherwise.
func (s *Set[Type]) Has(value Type) bool {
	_, ok := s.set[value]
	return ok
}

// Remove deletes the specified value from the set. It has no effect if the value does not exist in the set.
func (s *Set[Type]) Remove(value Type) {
	delete(s.set, value)
}

// Size returns the number of unique elements currently stored in the set.
func (s *Set[Type]) Size() int {
	return len(s.set)
}

// Union returns a new Set containing all unique elements from the current set and the specified other set.
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

// Intersection returns a new Set containing elements present in both the current set and the specified other set.
func (s *Set[Type]) Intersection(other *Set[Type]) *Set[Type] {
	intersection := New[Type]()
	for value := range s.set {
		if other.Has(value) {
			intersection.Add(value)
		}
	}
	return intersection
}

// Difference returns a new Set containing elements present in the current set but not in the specified other set.
func (s *Set[Type]) Difference(other *Set[Type]) *Set[Type] {
	difference := New[Type]()
	for value := range s.set {
		if !other.Has(value) {
			difference.Add(value)
		}
	}
	return difference
}

// SymmetricDifference returns a new Set containing elements that are in either the current set or the other set but not both.
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

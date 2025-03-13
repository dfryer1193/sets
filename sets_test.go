package sets

import (
	"testing"
)

func TestNew(t *testing.T) {
	set := New[int]()
	if set == nil {
		t.Error("New() returned nil")
	}
	if set.Size() != 0 {
		t.Errorf("New() set size = %d; want 0", set.Size())
	}
}

func TestFrom(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		wantSize int
	}{
		{"empty", []int{}, 0},
		{"single value", []int{1}, 1},
		{"multiple values", []int{1, 2, 3}, 3},
		{"duplicate values", []int{1, 1, 2, 2, 3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := From(tt.values...)
			if got := set.Size(); got != tt.wantSize {
				t.Errorf("From() size = %v, want %v", got, tt.wantSize)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	set := New[string]()
	
	// Test adding new elements
	set.Add("a")
	if !set.Has("a") {
		t.Error("Add() failed to add element")
	}
	
	// Test adding duplicate elements
	set.Add("a")
	if set.Size() != 1 {
		t.Error("Add() should not increase size for duplicate elements")
	}
}

func TestSet_Has(t *testing.T) {
	set := From("a", "b", "c")
	
	tests := []struct {
		value string
		want  bool
	}{
		{"a", true},
		{"b", true},
		{"d", false},
	}

	for _, tt := range tests {
		if got := set.Has(tt.value); got != tt.want {
			t.Errorf("Has(%v) = %v, want %v", tt.value, got, tt.want)
		}
	}
}

func TestSet_Remove(t *testing.T) {
	set := From(1, 2, 3)
	
	set.Remove(2)
	if set.Has(2) {
		t.Error("Remove() failed to remove element")
	}
	
	// Remove non-existent element
	set.Remove(4)
	if set.Size() != 2 {
		t.Error("Remove() of non-existent element changed set size")
	}
}

func TestSet_Size(t *testing.T) {
	set := New[int]()
	if set.Size() != 0 {
		t.Error("Size() of empty set should be 0")
	}

	set.Add(1)
	set.Add(2)
	if set.Size() != 2 {
		t.Error("Size() returned incorrect value")
	}
}

func TestSet_Union(t *testing.T) {
	set1 := From(1, 2, 3)
	set2 := From(3, 4, 5)
	
	union := set1.Union(set2)
	
	expectedSize := 5
	if union.Size() != expectedSize {
		t.Errorf("Union() size = %v, want %v", union.Size(), expectedSize)
	}
	
	for _, v := range []int{1, 2, 3, 4, 5} {
		if !union.Has(v) {
			t.Errorf("Union() missing element %v", v)
		}
	}
}

func TestSet_Intersection(t *testing.T) {
	set1 := From(1, 2, 3, 4)
	set2 := From(3, 4, 5, 6)
	
	intersection := set1.Intersection(set2)
	
	expectedSize := 2
	if intersection.Size() != expectedSize {
		t.Errorf("Intersection() size = %v, want %v", intersection.Size(), expectedSize)
	}
	
	for _, v := range []int{3, 4} {
		if !intersection.Has(v) {
			t.Errorf("Intersection() missing element %v", v)
		}
	}
}

func TestSet_Difference(t *testing.T) {
	set1 := From(1, 2, 3, 4)
	set2 := From(3, 4, 5, 6)
	
	difference := set1.Difference(set2)
	
	expectedSize := 2
	if difference.Size() != expectedSize {
		t.Errorf("Difference() size = %v, want %v", difference.Size(), expectedSize)
	}
	
	for _, v := range []int{1, 2} {
		if !difference.Has(v) {
			t.Errorf("Difference() missing element %v", v)
		}
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	set1 := From(1, 2, 3, 4)
	set2 := From(3, 4, 5, 6)
	
	symDiff := set1.SymmetricDifference(set2)
	
	expectedSize := 4
	if symDiff.Size() != expectedSize {
		t.Errorf("SymmetricDifference() size = %v, want %v", symDiff.Size(), expectedSize)
	}
	
	for _, v := range []int{1, 2, 5, 6} {
		if !symDiff.Has(v) {
			t.Errorf("SymmetricDifference() missing element %v", v)
		}
	}
}
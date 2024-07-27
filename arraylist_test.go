package collections

import (
	"testing"
)

// Test for GetIndex
func TestGetIndex(t *testing.T) {
	list := NewArrayListT(1, 2, 3, 4, 5)

	index := list.GetIndex(3, func(a, b int) bool { return a == b })
	if index != 2 {
		t.Errorf("Expected index 2, got %d", index)
	}

	index = list.GetIndex(10, func(a, b int) bool { return a == b })
	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

// Test for Remove
func TestRemove(t *testing.T) {
	list := NewArrayListT(1, 2, 3, 4, 5)

	removed := list.Remove(3, func(a, b int) bool { return a == b })
	if !removed {
		t.Error("Expected item to be removed, but it was not")
	}
	expected := []int{1, 2, 4, 5}
	for i, item := range list.Items() {
		if item != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], item)
		}
	}

	removed = list.Remove(10, func(a, b int) bool { return a == b })
	if removed {
		t.Error("Expected item to not be removed, but it was")
	}
}

func TestArrayListIterator(t *testing.T) {
	list := NewArrayListT(1, 2, 3, 4, 5)
	iterator := list.NewIterator()

	// Test HasNext and Next
	expectedValues := []int{1, 2, 3, 4, 5}
	for i, expected := range expectedValues {
		if !iterator.HasNext() {
			t.Errorf("Expected more elements, but HasNext() returned false at index %d", i)
		}
		item, ok := iterator.Next()
		if !ok {
			t.Errorf("Expected to get an item, but Next() returned false at index %d", i)
		}
		if item != expected {
			t.Errorf("Expected %d, got %d at index %d", expected, item, i)
		}
	}

	// Test HasNext when no more elements are left
	if iterator.HasNext() {
		t.Error("Expected HasNext() to return false, but it returned true when all elements were iterated")
	}

	// Test Next when no more elements are left
	item, ok := iterator.Next()
	if ok {
		t.Errorf("Expected Next() to return false when no more elements are left, but it returned %v", item)
	}

	// Test iterator with an empty list
	emptyList := NewArrayList[int]()
	emptyIterator := emptyList.NewIterator()

	if emptyIterator.HasNext() {
		t.Error("Expected HasNext() to return false for empty list, but it returned true")
	}

	item, ok = emptyIterator.Next()
	if ok {
		t.Errorf("Expected Next() to return false for empty list, but it returned %v", item)
	}
}

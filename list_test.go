package collections

import "testing"

// TestNewList tests the NewList function.
func TestNewList(t *testing.T) {
	list := NewList[int]()
	if list.Count() != 0 {
		t.Errorf("Expected empty list, got count %d", list.Count())
	}
}

// TestNewListT tests the NewListT function with initial items.
func TestNewListT(t *testing.T) {
	list := NewListT(1, 2, 3, 4, 5)
	if list.Count() != 5 {
		t.Errorf("Expected list with 5 items, got count %d", list.Count())
	}
	if items := list.Items(); len(items) != 5 {
		t.Errorf("Expected 5 items, got %d", len(items))
	}
}

// TestAdd tests the Add method.
func TestAdd(t *testing.T) {
	list := NewList[int]()
	list.Add(1)
	list.Add(2)
	if list.Count() != 2 {
		t.Errorf("Expected list with 2 items, got count %d", list.Count())
	}
	if items := list.Items(); len(items) != 2 || items[0] != 1 || items[1] != 2 {
		t.Errorf("Expected items [1 2], got %v", items)
	}
}

// TestAddRange tests the AddRange method.
func TestAddRange(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3})
	if list.Count() != 3 {
		t.Errorf("Expected list with 3 items, got count %d", list.Count())
	}
	if items := list.Items(); len(items) != 3 || items[0] != 1 || items[1] != 2 || items[2] != 3 {
		t.Errorf("Expected items [1 2 3], got %v", items)
	}
}

// TestGet tests the Get method.
func TestGet(t *testing.T) {
	list := NewListT(1, 2, 3)
	if item := list.Get(1); item != 2 {
		t.Errorf("Expected item 2, got %d", item)
	}
}

// TestGetIndex tests the GetIndex method.
func TestGetIndex(t *testing.T) {
	list := NewListT(1, 2, 3)
	if index := list.GetIndex(2); index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}
	if index := list.GetIndex(4); index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

// TestRemove tests the Remove method.
func TestRemove(t *testing.T) {
	list := NewListT(1, 2, 3)
	if removed := list.Remove(2); !removed {
		t.Error("Expected item 2 to be removed")
	}
	if list.Count() != 2 {
		t.Errorf("Expected list with 2 items, got count %d", list.Count())
	}
	if items := list.Items(); len(items) != 2 || items[0] != 1 || items[1] != 3 {
		t.Errorf("Expected items [1 3], got %v", items)
	}
}

// TestRemoveAt tests the RemoveAt method.
func TestRemoveAt(t *testing.T) {
	list := NewListT(1, 2, 3)
	list.RemoveAt(1)
	if list.Count() != 2 {
		t.Errorf("Expected list with 2 items, got count %d", list.Count())
	}
	if items := list.Items(); len(items) != 2 || items[0] != 1 || items[1] != 3 {
		t.Errorf("Expected items [1 3], got %v", items)
	}
}

// TestFilter tests the Filter method.
func TestFilter(t *testing.T) {
	list := NewListT(1, 2, 3, 4, 5)
	evenNumbers := list.Filter(func(item int) bool { return item%2 == 0 })
	if len(evenNumbers) != 2 || evenNumbers[0] != 2 || evenNumbers[1] != 4 {
		t.Errorf("Expected even numbers [2 4], got %v", evenNumbers)
	}
}

// TestFirstOrDefault tests the FirstOrDefault method.
func TestFirstOrDefault(t *testing.T) {
	list := NewListT(1, 2, 3, 4, 5)
	firstEven := list.FirstOrDefault(func(item int) bool { return item%2 == 0 })
	if firstEven != 2 {
		t.Errorf("Expected first even number 2, got %d", firstEven)
	}
}

// TestLastOrDefault tests the LastOrDefault method.
func TestLastOrDefault(t *testing.T) {
	list := NewListT(1, 2, 3, 4, 5)
	lastEven := list.LastOrDefault(func(item int) bool { return item%2 == 0 })
	if lastEven != 4 {
		t.Errorf("Expected last even number 4, got %d", lastEven)
	}
}

// TestSingleOrDefault tests the SingleOrDefault method.
func TestSingleOrDefault(t *testing.T) {
	list := NewListT(1, 2, 3)
	singleEven := list.SingleOrDefault(func(item int) bool { return item == 2 })
	if singleEven != 2 {
		t.Errorf("Expected single even number 2, got %d", singleEven)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for more than one matching element, but did not panic")
		}
	}()
	list.SingleOrDefault(func(item int) bool { return item < 4 })
}

// TestSet tests the Set method.
func TestSet(t *testing.T) {
	list := NewListT(1, 2, 3)
	list.Set(1, 10)
	if item := list.Get(1); item != 10 {
		t.Errorf("Expected item 10 at index 1, got %d", item)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for index out of range, but did not panic")
		}
	}()
	list.Set(10, 100)
}

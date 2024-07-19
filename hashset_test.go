package collections

import "testing"

func TestHashSet(t *testing.T) {
    set := NewHashSet[string]()

    // Test Add
    if added := set.Add("item1"); !added {
        t.Errorf("Add() = false; want true")
    }
    if added := set.Add("item1"); added {
        t.Errorf("Add() = true; want false")
    }

    // Test Contains
    if !set.Contains("item1") {
        t.Errorf("Contains() = false; want true")
    }
    if set.Contains("item2") {
        t.Errorf("Contains() = true; want false")
    }

    // Test Remove
    if removed := set.Remove("item1"); !removed {
        t.Errorf("Remove() = false; want true")
    }
    if removed := set.Remove("item1"); removed {
        t.Errorf("Remove() = true; want false")
    }

    // Test Count
    set.Add("item2")
    set.Add("item3")
    if count := set.Count(); count != 2 {
        t.Errorf("Count() = %d; want 2", count)
    }

    // Test Clear
    set.Clear()
    if count := set.Count(); count != 0 {
        t.Errorf("Count() after Clear() = %d; want 0", count)
    }

    // Test Items
    set.Add("item4")
    set.Add("item5")
    items := set.Items()
    if len(items) != 2 {
        t.Errorf("Items() = %v; want length 2", items)
    }
}

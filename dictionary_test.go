package collections

import "testing"

func TestDictionary(t *testing.T) {
    dict := NewDictionary[string, int]()

    // Test Set and Get
    dict.Set("key1", 1)
    if value, ok := dict.Get("key1"); !ok || value != 1 {
        t.Errorf("Get() = %v, %v; want 1, true", value, ok)
    }

    // Test Remove
    dict.Remove("key1")
    if _, ok := dict.Get("key1"); ok {
        t.Errorf("Get() after Remove() should be false")
    }

    // Test Keys and Values
    dict.Set("key2", 2)
    dict.Set("key3", 3)
    keys := dict.Keys()
    values := dict.Values()

    if len(keys) != 2 || len(values) != 2 {
        t.Errorf("Keys() or Values() have unexpected length")
    }

    // Test Count
    if count := dict.Count(); count != 2 {
        t.Errorf("Count() = %d; want 2", count)
    }
}

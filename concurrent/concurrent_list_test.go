package concurrent

import (
	"sync"
	"testing"
)

func TestConcurrentList(t *testing.T) {
    list := NewConcurrentList[int]()
    var wg sync.WaitGroup

    // Test concurrent additions
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            list.Add(i)
        }(i)
    }

    wg.Wait()

    if list.Count() != 100 {
        t.Errorf("Expected 100 items, but got %d", list.Count())
    }
}

func TestConcurrentListModification(t *testing.T) {
    list := NewConcurrentListT(1, 2, 3, 4, 5)
    var wg sync.WaitGroup

    // Test concurrent modifications
    wg.Add(2)
    go func() {
        defer wg.Done()
        list.Remove(2)
    }()
    go func() {
        defer wg.Done()
        list.Add(6)
    }()

    wg.Wait()

    expected := []int{1, 3, 4, 5, 6}
    result := list.Items()

    if len(result) != len(expected) {
        t.Errorf("Expected %v but got %v", expected, result)
        return
    }

    for i, v := range expected {
        if result[i] != v {
            t.Errorf("Expected %d but got %d at index %d", v, result[i], i)
        }
    }
}

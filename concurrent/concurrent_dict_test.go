package concurrent

import (
	"sync"
	"testing"
)

func TestConcurrentDict(t *testing.T) {
    dict := NewConcurrentDict[string, int]()
    var wg sync.WaitGroup

    // Test concurrent operations
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            dict.Set(string(i), i)
        }(i)
    }

    wg.Wait()

    // Verify all keys are present
    for i := 0; i < 100; i++ {
        value, ok := dict.Get(string(i))
        if !ok || value != i {
            t.Errorf("Expected %d for key %d but got %v", i, i, value)
        }
    }
}

func TestConcurrentDictDeletion(t *testing.T) {
    dict := NewConcurrentDict[string, int]()
    for i := 0; i < 100; i++ {
        dict.Set(string(i), i)
    }

    var wg sync.WaitGroup

    // Test concurrent deletions
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            dict.Delete(string(i))
        }(i)
    }

    wg.Wait()

    // Verify that some keys are deleted
    for i := 0; i < 50; i++ {
        _, ok := dict.Get(string(i))
        if ok {
            t.Errorf("Key %d was expected to be deleted but was found", i)
        }
    }

    // Verify remaining keys
    for i := 50; i < 100; i++ {
        value, ok := dict.Get(string(i))
        if !ok || value != i {
            t.Errorf("Expected %d for key %d but got %v", i, i, value)
        }
    }
}

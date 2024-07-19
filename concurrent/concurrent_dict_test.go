package concurrent

import (
	"strconv"
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
            key := strconv.Itoa(i) // Use strconv to convert int to string
            dict.Set(key, i)
        }(i)
    }

    wg.Wait()

    // Verify all keys are present
    for i := 0; i < 100; i++ {
        key := strconv.Itoa(i) // Use strconv to convert int to string
        value, ok := dict.Get(key)
        if !ok || value != i {
            t.Errorf("Expected %d for key %d but got %v", i, i, value)
        }
    }
}

func TestConcurrentDictDeletion(t *testing.T) {
    dict := NewConcurrentDict[string, int]()
    for i := 0; i < 100; i++ {
        key := strconv.Itoa(i) // Use strconv to convert int to string
        dict.Set(key, i)
    }

    var wg sync.WaitGroup

    // Test concurrent deletions
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            key := strconv.Itoa(i) // Use strconv to convert int to string
            dict.Delete(key)
        }(i)
    }

    wg.Wait()

    // Verify that some keys are deleted
    for i := 0; i < 50; i++ {
        key := strconv.Itoa(i) // Use strconv to convert int to string
        _, ok := dict.Get(key)
        if ok {
            t.Errorf("Key %d was expected to be deleted but was found", i)
        }
    }

    // Verify remaining keys
    for i := 50; i < 100; i++ {
        key := strconv.Itoa(i) // Use strconv to convert int to string
        value, ok := dict.Get(key)
        if !ok || value != i {
            t.Errorf("Expected %d for key %d but got %v", i, i, value)
        }
    }
}

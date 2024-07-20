package concurrent

import (
	"sort"
	"sync"
)

// ConcurrentList is a thread-safe list.
type ConcurrentList[T comparable] struct {
	mu         sync.Mutex
	collection []T
}

// NewConcurrentList initializes a new empty ConcurrentList.
func NewConcurrentList[T comparable]() *ConcurrentList[T] {
	return &ConcurrentList[T]{collection: []T{}}
}

// NewConcurrentListT initializes a new ConcurrentList with the given items.
func NewConcurrentListT[T comparable](items ...T) *ConcurrentList[T] {
	l := &ConcurrentList[T]{collection: make([]T, len(items))}
	copy(l.collection, items)
	return l
}

// Count returns the number of items in the ConcurrentList.
func (l *ConcurrentList[T]) Count() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.collection)
}

// Clear removes all items from the ConcurrentList.
func (l *ConcurrentList[T]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.collection = []T{}
}

// Items returns a slice of all items in the ConcurrentList.
func (l *ConcurrentList[T]) Items() []T {
	l.mu.Lock()
	defer l.mu.Unlock()
	// Return a copy of the slice to avoid data races
	result := make([]T, len(l.collection))
	copy(result, l.collection)
	return result
}

// Add appends an item to the ConcurrentList.
func (l *ConcurrentList[T]) Add(item T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.collection = append(l.collection, item)
}

// AddRange appends multiple items to the ConcurrentList.
func (l *ConcurrentList[T]) AddRange(items []T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.collection = append(l.collection, items...)
}

// Get retrieves the item at the specified index.
func (l *ConcurrentList[T]) Get(index int) T {
	l.mu.Lock()
	defer l.mu.Unlock()
	if index < 0 || index >= len(l.collection) {
		panic("Index out of range.")
	}
	return l.collection[index]
}

// Set updates the item at the specified index.
func (l *ConcurrentList[T]) Set(index int, item T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if index < 0 || index >= len(l.collection) {
		panic("Index out of range.")
	}
	l.collection[index] = item
}

// Remove removes the first occurrence of the specified item from the ConcurrentList.
func (l *ConcurrentList[T]) Remove(item T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, v := range l.collection {
		if v == item {
			l.collection = append(l.collection[:i], l.collection[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveAt removes the item at the specified index.
func (l *ConcurrentList[T]) RemoveAt(index int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if index < 0 || index >= len(l.collection) {
		panic("Index out of range.")
	}
	l.collection = append(l.collection[:index], l.collection[index+1:]...)
}

// OrderBy sorts the list using the provided less function
func (l *ConcurrentList[T]) OrderBy(less func(i, j T) bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	sort.Slice(l.collection, func(i, j int) bool {
		return less(l.collection[i], l.collection[j])
	})
}

// Filter returns a new slice containing all items that match the predicate
func (l *ConcurrentList[T]) Filter(predicate func(T) bool) []T {
	l.mu.Lock()
	defer l.mu.Unlock()

	results := make([]T, 0)
	for _, val := range l.collection {
		if predicate(val) {
			results = append(results, val)
		}
	}

	return results
}

// FirstOrDefault returns the first item that matches the predicate, or the zero value of T if no match is found.
// This method locks the list for thread safety.
func (l *ConcurrentList[T]) FirstOrDefault(predicate func(T) bool) T {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, val := range l.collection {
		if predicate(val) {
			return val
		}
	}

	var zeroValue T
	return zeroValue
}

// LastOrDefault returns the last item that matches the predicate, or the zero value of T if no match is found.
// This method locks the list for thread safety.
func (l *ConcurrentList[T]) LastOrDefault(predicate func(T) bool) T {
	l.mu.Lock()
	defer l.mu.Unlock()

	var zeroValue T
	for _, val := range l.collection {
		if predicate(val) {
			zeroValue = val
		}
	}

	return zeroValue
}

// SingleOrDefault returns the single item that matches the predicate, or the zero value of T if no match is found.
// If more than one item matches, it panics. This method locks the list for thread safety.
func (l *ConcurrentList[T]) SingleOrDefault(predicate func(T) bool) T {
	l.mu.Lock()
	defer l.mu.Unlock()

	var result T
	count := 0
	for _, val := range l.collection {
		if predicate(val) {
			if count == 0 {
				result = val
				count++
			} else {
				panic("Sequence contains more than one matching element.")
			}
		}
	}

	return result
}

// ConcurrentListIterator is an iterator for ConcurrentList.
type ConcurrentListIterator[T comparable] struct {
	list  *ConcurrentList[T]
	index int
	mu    sync.Mutex
}

// NewConcurrentListIterator creates a new iterator for the ConcurrentList.
func (l *ConcurrentList[T]) NewIterator() *ConcurrentListIterator[T] {
	return &ConcurrentListIterator[T]{list: l, index: -1}
}

// HasNext returns true if there are more items to iterate over.
func (it *ConcurrentListIterator[T]) HasNext() bool {
	it.mu.Lock()
	defer it.mu.Unlock()

	it.list.mu.Lock()
	defer it.list.mu.Unlock()

	return it.index+1 < len(it.list.collection)
}

// Next returns the next item in the iteration.
func (it *ConcurrentListIterator[T]) Next() T {
	it.mu.Lock()
	defer it.mu.Unlock()

	it.list.mu.Lock()
	defer it.list.mu.Unlock()

	if !it.HasNext() {
		var zeroValue T
		return zeroValue // or you might choose to panic or return an error
	}

	it.index++
	return it.list.collection[it.index]
}

package concurrent

import (
	"sync"
)

// ConcurrentQueue represents a generic, thread-safe queue data structure.
type ConcurrentQueue[T any] struct {
	mu       sync.Mutex
	elements []T
}

// NewConcurrentQueue creates a new instance of a ConcurrentQueue.
func NewConcurrentQueue[T any]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{}
}

// Enqueue adds an element to the end of the queue.
func (q *ConcurrentQueue[T]) Enqueue(element T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.elements = append(q.elements, element)
}

// Dequeue removes and returns the element from the front of the queue.
// It panics if the queue is empty.
func (q *ConcurrentQueue[T]) Dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.elements) == 0 {
		panic("Dequeue from an empty queue")
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

// Peek returns the element at the front of the queue without removing it.
// It panics if the queue is empty.
func (q *ConcurrentQueue[T]) Peek() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.elements) == 0 {
		panic("Peek from an empty queue")
	}

	return q.elements[0]
}

// IsEmpty returns true if the queue is empty, false otherwise.
func (q *ConcurrentQueue[T]) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.elements) == 0
}

// Size returns the number of elements in the queue.
func (q *ConcurrentQueue[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.elements)
}

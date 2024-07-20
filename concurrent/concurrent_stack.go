package concurrent

import "sync"

// ConcurrentStack represents a generic stack data structure.
type ConcurrentStack[T any] struct {
	mu       sync.Mutex
	elements []T
}

// NewConcurrentStack creates a new instance of a Stack.
func NewConcurrentStack[T any]() *ConcurrentStack[T] {
	return &ConcurrentStack[T]{}
}

// Push adds an element to the top of the ConcurrentStack.
func (s *ConcurrentStack[T]) Push(element T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element of the ConcurrentStack.
// It panics if the stack is empty.
func (s *ConcurrentStack[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.elements) == 0 {
		panic("Pop from an empty stack")
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

// Peek returns the top element of the ConcurrentStack without removing it.
// It panics if the stack is empty.
func (s *ConcurrentStack[T]) Peek() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.elements) == 0 {
		panic("Peek from an empty stack")
	}

	return s.elements[len(s.elements)-1]
}

// IsEmpty returns true if the ConcurrentStack is empty, false otherwise.
func (s *ConcurrentStack[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.elements) == 0
}

// Size returns the number of elements in the ConcurrentStack.
func (s *ConcurrentStack[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.elements)
}

package collections

type List[T any] struct {
	collection []T
}

// NewList initializes a new empty List
//
// Example:
//  list := NewList[int]()
//  fmt.Println(list) // Output: &{[]}
func NewList[T any]() *List[T] {
	return &List[T]{
		collection: []T{},
	}
}

// NewListT initializes a new List with the given items
//
// Example:
//  list := NewListT(1, 2, 3, 4, 5, 6)
//  fmt.Println(list) // Output: &{[1 2 3 4 5 6]}
func NewListT[T any](items ...T) *List[T] {
	l := List[T]{
		collection: make([]T, len(items)),
	}

	// Copy the items into the collection
	copy(l.collection, items)

	return &l
}

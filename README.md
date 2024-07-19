
# Collections Package

The `collections` package provides a generic `List` type in Go that holds a collection of items of any comparable type. This package includes methods for managing the list, including adding, removing, and accessing items, as well as querying the list with predicates.

## Features

- **Generic List:** Holds a collection of items of any comparable type.
- **Add/Remove Items:** Methods to add and remove items.
- **Access Items:** Retrieve items by index or value.
- **Filter and Query:** Methods to filter and query items based on predicates.
- **Indexer:** Provides indexing functionality for direct access to items.

## Installation

To install the `collections` package, include it in your Go module:

```bash
go get github.com/VikashChauhan51/collections
```

## Usage

### Creating a New List

```go
package main

import (
    "fmt"
    "github.com/VikashChauhan51/collections"
)

func main() {
    list := collections.NewList[int]()
    fmt.Println(list) // Output: &{[]}

    list.Add(1)
    list.AddRange([]int{2, 3})
    fmt.Println(list) // Output: &{[1 2 3]}

    // Access items using Indexer
    fmt.Println(list.Index(1)) // Output: 2

    // Remove an item
    list.Remove(2)
    fmt.Println(list) // Output: &{[1 3]}
}
```

### Functions and Methods

#### `NewList[T comparable]() *List[T]`

Initializes a new empty List.

**Example:**

```go
list := collections.NewList[int]()
fmt.Println(list) // Output: &{[]}
```

#### `NewListT[T comparable](items ...T) *List[T]`

Initializes a new List with the given items.

**Example:**

```go
list := collections.NewListT(1, 2, 3, 4, 5, 6)
fmt.Println(list) // Output: &{[1 2 3 4 5 6]}
```

#### `Count() int`

Returns the number of items in the List.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
fmt.Println(list.Count()) // Output: 3
```

#### `Clear()`

Removes all items from the List.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
list.Clear()
fmt.Println(list) // Output: &{[]}
```

#### `Items() []T`

Returns a slice of all items in the List.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
items := list.Items()
fmt.Println(items) // Output: [1 2 3]
```

#### `Add(item T)`

Appends an item to the List.

**Example:**

```go
list := collections.NewList[int]()
list.Add(1)
fmt.Println(list) // Output: &{[1]}
```

#### `AddRange(items []T)`

Appends multiple items to the List.

**Example:**

```go
list := collections.NewList[int]()
list.AddRange([]int{1, 2, 3})
fmt.Println(list) // Output: &{[1 2 3]}
```

#### `Get(index int) T`

Retrieves the item at the specified index.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
item := list.Get(1)
fmt.Println(item) // Output: 2
```

#### `Set(index int, item T)`

Updates the item at the specified index.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
list.Set(1, 10)
fmt.Println(list) // Output: &{[1 10 3]}
```

#### `GetIndex(item T) int`

Retrieves the index of the specified item in the List. Returns -1 if the item is not found.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
index := list.GetIndex(2)
fmt.Println(index) // Output: 1
```

#### `Remove(item T) bool`

Removes the first occurrence of the specified item from the List.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
list.Remove(2)
fmt.Println(list) // Output: &{[1 3]}
```

#### `RemoveAt(index int)`

Removes the item at the specified index.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
list.RemoveAt(1)
fmt.Println(list) // Output: &{[1 3]}
```

#### `Filter(predicate func(T) bool) []T`

Returns a new slice containing all items that match the predicate.

**Example:**

```go
list := collections.NewListT(1, 2, 3, 4, 5)
evenNumbers := list.Filter(func(item int) bool { return item%2 == 0 })
fmt.Println(evenNumbers) // Output: [2 4]
```

#### `FirstOrDefault(predicate func(T) bool) T`

Returns the first item that matches the predicate. If no item matches, it returns the zero value of the type.

**Example:**

```go
list := collections.NewListT(1, 2, 3, 4, 5)
firstEven := list.FirstOrDefault(func(item int) bool { return item%2 == 0 })
fmt.Println(firstEven) // Output: 2
```

#### `LastOrDefault(predicate func(T) bool) T`

Returns the last item that matches the predicate. If no item matches, it returns the zero value of the type.

**Example:**

```go
list := collections.NewListT(1, 2, 3, 4, 5)
lastEven := list.LastOrDefault(func(item int) bool { return item%2 == 0 })
fmt.Println(lastEven) // Output: 4
```

#### `SingleOrDefault(predicate func(T) bool) T`

Returns the single item that matches the predicate. If no item or more than one item matches, it panics.

**Example:**

```go
list := collections.NewListT(1, 2, 3, 4, 5)
singleEven := list.SingleOrDefault(func(item int) bool { return item == 2 })
fmt.Println(singleEven) // Output: 2
```

#### `Index(index int) T`

Provides indexing-like access to items in the List. This method allows you to use syntax similar to `list[index]`.

**Example:**

```go
list := collections.NewListT(1, 2, 3)
item := list.Index(1)
fmt.Println(item) // Output: 2
```

## Contributing

If you would like to contribute to this package, please fork the repository and submit a pull request. Ensure that your code passes all tests and follows the project's coding style.

## License

This package is licensed under the [MIT License](LICENSE).

---

Feel free to modify the paths and examples according to your actual module and usage scenarios.
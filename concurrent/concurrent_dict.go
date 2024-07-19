package concurrent

import "sync"

// ConcurrentDict is a thread-safe dictionary.
type ConcurrentDict[K comparable, V any] struct {
    m sync.Map
}

// NewConcurrentDict initializes a new empty ConcurrentDict.
func NewConcurrentDict[K comparable, V any]() *ConcurrentDict[K, V] {
    return &ConcurrentDict[K, V]{}
}

// Set adds or updates the value for the given key.
func (d *ConcurrentDict[K, V]) Set(key K, value V) {
    d.m.Store(key, value)
}

// Get retrieves the value for the given key.
func (d *ConcurrentDict[K, V]) Get(key K) (V, bool) {
    value, ok := d.m.Load(key)
    if ok {
        return value.(V), true
    }
    var zeroValue V
    return zeroValue, false
}

// Delete removes the value for the given key.
func (d *ConcurrentDict[K, V]) Delete(key K) {
    d.m.Delete(key)
}

// Keys returns a slice of all keys in the ConcurrentDict.
func (d *ConcurrentDict[K, V]) Keys() []K {
    var keys []K
    d.m.Range(func(key, _ interface{}) bool {
        keys = append(keys, key.(K))
        return true
    })
    return keys
}

// Values returns a slice of all values in the ConcurrentDict.
func (d *ConcurrentDict[K, V]) Values() []V {
    var values []V
    d.m.Range(func(_, value interface{}) bool {
        values = append(values, value.(V))
        return true
    })
    return values
}

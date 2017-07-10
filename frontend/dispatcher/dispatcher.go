// Package dispatcher is a simple flux-like implementation
package dispatcher

// ID is returned when you register a dispatcher callback. Use the ID to unregister your callback.
type ID int

var id ID
var callbacks = map[ID]func(){}

// Register takes a function and returns an ID. Use the ID to unregister your callback.
func Register(f func()) ID {
	id++
	callbacks[id] = f
	return id
}

// Dispatch calls every registered function
func Dispatch() {
	for _, cb := range callbacks {
		cb()
	}
}

// Unregister removes the callback that belongs to the given ID.
func Unregister(id ID) {
	delete(callbacks, id)
}

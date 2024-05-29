package set

// Set interface represents simple set ADT.
type Set[T comparable] interface {
	// Add ads an element to this Set. If this element is already present,
	// then Add is no-op.
	Add(element T)

	// Contains returns true iff this Set contains element.
	Contains(element T) bool

	// Elements return a slice of all elements in this Set with unspecified order.
	//
	// No operation on the returned slice may affect the state of this Set.
	Elements() []T

	// Size returns current Set size.
	Size() int

	// IsEmpty returns true iff Size() == 0.
	IsEmpty() bool
}

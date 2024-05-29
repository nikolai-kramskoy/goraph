package queue

// Queue interface represents simple queue ADT.
type Queue[T any] interface {
	// Push adds an element to the tail of the Queue.
	Push(element T)

	// Pop removes an element from the head of the Queue.
	//
	// Panics if one try to Pop from empty Queue.
	Pop() T

	// Size returns current Queue size.
	Size() int

	// IsEmpty returns true iff Size() == 0.
	IsEmpty() bool
}

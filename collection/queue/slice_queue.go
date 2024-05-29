package queue

type sliceQueue[T any] struct {
	slice []T

	// We take an element from queue's head
	headIndex int

	// We add an element to queue's tail
	tailIndex int
	size      int
}

var _ Queue[struct{}] = (*sliceQueue[struct{}])(nil)

// NewEmptySliceQueue creates an empty slice-based Queue.
//
// It is a non-shrinking slice-based Queue, be aware that with time this
// data structure will consume a lot of memory.
func NewEmptySliceQueue[T any]() Queue[T] {
	slice := make([]T, 0)

	return &sliceQueue[T]{
		slice,
		-1,
		-1,
		0,
	}
}

func (queue *sliceQueue[T]) Push(element T) {
	queue.slice = append(queue.slice, element)

	if queue.size == 0 {
		queue.headIndex++
	}

	queue.tailIndex++
	queue.size++
}

func (queue *sliceQueue[T]) Pop() T {
	if queue.size == 0 {
		panic("Pop from an empty queue")
	}

	elementIndex := queue.headIndex

	if queue.size > 1 {
		queue.headIndex++
	}

	queue.size--

	return queue.slice[elementIndex]
}

func (queue *sliceQueue[T]) Size() int {
	return queue.size
}

func (queue *sliceQueue[T]) IsEmpty() bool {
	return queue.Size() == 0
}

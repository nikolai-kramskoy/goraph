package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceQueue_1(t *testing.T) {
	intQueue := NewEmptySliceQueue[int]()
	assert.NotNil(t, intQueue)
	assertSize(t, 0, intQueue)
	assert.Panics(t, func() { intQueue.Pop() })

	intQueue.Push(50)
	assertSize(t, 1, intQueue)

	intQueue.Push(100)
	assertSize(t, 2, intQueue)

	assert.Equal(t, 50, intQueue.Pop())
	assertSize(t, 1, intQueue)

	assert.Equal(t, 100, intQueue.Pop())
	assertSize(t, 0, intQueue)
	assert.Panics(t, func() { intQueue.Pop() })
}

func TestSliceQueue_2(t *testing.T) {
	intQueue := NewEmptySliceQueue[int]()
	assert.NotNil(t, intQueue)
	assertSize(t, 0, intQueue)
	assert.Panics(t, func() { intQueue.Pop() })

	intQueue.Push(50)
	assertSize(t, 1, intQueue)

	assert.Equal(t, 50, intQueue.Pop())
	assertSize(t, 0, intQueue)
	assert.Panics(t, func() { intQueue.Pop() })

	intQueue.Push(100)
	assertSize(t, 1, intQueue)

	assert.Equal(t, 100, intQueue.Pop())
	assertSize(t, 0, intQueue)
	assert.Panics(t, func() { intQueue.Pop() })
}

func assertSize(t *testing.T, expectedSize int, intQueue Queue[int]) {
	assert.Equal(t, expectedSize, intQueue.Size())

	if expectedSize == 0 {
		assert.True(t, intQueue.IsEmpty())
	} else {
		assert.False(t, intQueue.IsEmpty())
	}
}

package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSet_1(t *testing.T) {
	intSet := NewEmptyMapSet[int]()
	assert.NotNil(t, intSet)
	assertSize(t, 0, intSet)

	intSet.Add(5)
	assertSize(t, 1, intSet)

	intSet.Add(5)
	assertSize(t, 1, intSet)

	intSet.Add(1)
	assertSize(t, 2, intSet)
}

func TestMapSet_Contains(t *testing.T) {
	intSet := NewEmptyMapSet[int]()
	assert.NotNil(t, intSet)
	assert.False(t, intSet.Contains(2))

	intSet.Add(2)
	assert.True(t, intSet.Contains(2))
}

func TestNewEmptyMapSet(t *testing.T) {
	intSet := NewEmptyMapSet[int]()
	assert.NotNil(t, intSet)
	assert.Equal(t, []int{}, intSet.Elements())

	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(2)
	intSet.Add(2)
	intSet.Add(3)
	intSet.Add(-100)
	intSet.Add(0)

	assertMapSetElements(t, intSet)
}

func TestNewMapSet(t *testing.T) {
	intSet := NewMapSet(1, 2, 2, 2, 3, -100, 0)
	assert.NotNil(t, intSet)

	assertMapSetElements(t, intSet)
}

func TestNewMapSetFromSlice(t *testing.T) {
	intSet := NewMapSetFromSlice([]int{1, 2, 2, 2, 3, -100, 0})
	assert.NotNil(t, intSet)

	assertMapSetElements(t, intSet)
}

func assertMapSetElements(t *testing.T, intSet Set[int]) {
	elements := intSet.Elements()
	assert.Equal(t, 5, len(elements))

	assert.Contains(t, elements, 1)
	assert.Contains(t, elements, 2)
	assert.Contains(t, elements, 3)
	assert.Contains(t, elements, -100)
	assert.Contains(t, elements, 0)
}

func assertSize(t *testing.T, expectedSize int, intSet Set[int]) {
	assert.Equal(t, expectedSize, intSet.Size())

	if expectedSize == 0 {
		assert.True(t, intSet.IsEmpty())
	} else {
		assert.False(t, intSet.IsEmpty())
	}
}

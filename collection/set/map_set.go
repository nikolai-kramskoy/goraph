package set

type mapSet[T comparable] struct {
	tToEmptyStruct map[T]struct{}
}

var _ Set[struct{}] = (*mapSet[struct{}])(nil)

// NewEmptyMapSet creates an empty map-based Set.
//
// As it is a map-based Set, it has inherent pros and cons of this data type,
// so be aware of possible pitfalls of comparing pointers.
func NewEmptyMapSet[T comparable]() Set[T] {
	return &mapSet[T]{map[T]struct{}{}}
}

// NewMapSet creates a map-based Set with non-duplicated entries from elements.
//
// Check NewEmptyMapSet doc for detailed explanation of this Set implementation.
func NewMapSet[T comparable](elements ...T) Set[T] {
	if elements == nil {
		panic("elements == nil")
	}

	tToEmptyStruct := make(map[T]struct{}, len(elements))

	for _, elem := range elements {
		tToEmptyStruct[elem] = struct{}{}
	}

	return &mapSet[T]{tToEmptyStruct}
}

// NewMapSetFromSlice creates a map-based Set with non-duplicated entries from elements slice.
//
// Check NewEmptyMapSet doc for detailed explanation of this Set implementation.
func NewMapSetFromSlice[T comparable](elements []T) Set[T] {
	if elements == nil {
		panic("elements == nil")
	}

	tToEmptyStruct := make(map[T]struct{}, len(elements))

	for _, elem := range elements {
		tToEmptyStruct[elem] = struct{}{}
	}

	return &mapSet[T]{tToEmptyStruct}
}

func (set *mapSet[T]) Add(element T) {
	set.tToEmptyStruct[element] = struct{}{}
}

func (set *mapSet[T]) Contains(element T) bool {
	_, isPresent := set.tToEmptyStruct[element]

	return isPresent
}

func (set *mapSet[T]) Elements() []T {
	elements := make([]T, len(set.tToEmptyStruct))

	i := 0
	for element := range set.tToEmptyStruct {
		elements[i] = element
		i++
	}

	return elements
}

func (set *mapSet[T]) Size() int {
	return len(set.tToEmptyStruct)
}

func (set *mapSet[T]) IsEmpty() bool {
	return set.Size() == 0
}

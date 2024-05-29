package graph

// Edge struct represents an immutable (so it is thread-safe) edge in Digraph.
type Edge[V Vertex] struct {
	source V
	target V
}

// NewEdge creates an immutable (so it is thread-safe) Edge.
func NewEdge[V Vertex](source, target V) Edge[V] {
	return Edge[V]{source, target}
}

// Source returns Source of this Edge.
func (edge *Edge[V]) Source() V {
	return edge.source
}

// Target returns Target of this Edge.
func (edge *Edge[V]) Target() V {
	return edge.target
}

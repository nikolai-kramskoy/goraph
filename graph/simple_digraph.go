package graph

// SimpleDigraph interface extends Digraph interface and is intended
// to represent immutable simple (i.e. without loops and without
// multiple edges) digraphs.
type SimpleDigraph[V Vertex] interface {
	Digraph[V]

	// Edge returns a non-nil *Edge iff edge with specified source and target
	// exists in this SimpleDigraph.
	//
	// No operation on the returned *Edge may affect the state of this SimpleDigraph.
	Edge(source, target V) *Edge[V]
}

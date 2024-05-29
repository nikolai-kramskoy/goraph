package graph

import "goraph/collection/set"

// Digraph interface represents an immutable digraph.
//
// This interface is designed to be extended so one may
// implement simple digraphs, multidigraphs etc.
type Digraph[V Vertex] interface {
	// Vertices returns a set.Set of all vertices in this Digraph.
	//
	// No operation on the returned set.Set may affect the state of this Digraph.
	Vertices() set.Set[V]

	// Edges returns a set.Set of all edges in this Digraph.
	//
	// No operation on the returned set.Set may affect the state of this Digraph.
	Edges() set.Set[Edge[V]]

	// Successors returns a set.Set of all vertices for which exists
	// at least one edge with vertex source and some Vertex target in this set.Set.
	//
	// No operation on the returned set.Set may affect the state of this Digraph.
	Successors(vertex V) set.Set[V]

	// Predecessors returns a set.Set of all vertices for which exists
	// at least one edge with vertex target and some Vertex source in this set.Set.
	//
	// No operation on the returned set.Set may affect the state of this Digraph.
	Predecessors(vertex V) set.Set[V]
}

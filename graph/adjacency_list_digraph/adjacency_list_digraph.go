package adjacency_list_digraph

import (
	"errors"
	"fmt"
	"goraph/collection/set"
	"goraph/graph"
)

type adjacencyListSimpleDigraph[V graph.Vertex] struct {
	successors   map[V]set.Set[V]
	predecessors map[V]set.Set[V]
}

var _ graph.SimpleDigraph[struct{}] = (*adjacencyListSimpleDigraph[struct{}])(nil)

// NewAdjacencyListSimpleDigraph creates an immutable graph.SimpleDigraph
// implementation using adjacency list ADT.
//
// This implementation is immutable and thread-safe.
//
// https://en.wikipedia.org/wiki/Adjacency_list
func NewAdjacencyListSimpleDigraph[V graph.Vertex](
	vertices set.Set[V],
	edges set.Set[graph.Edge[V]],
) (graph.SimpleDigraph[V], error) {
	if vertices == nil {
		return nil, errors.New("vertices == nil")
	}
	if edges == nil {
		return nil, errors.New("edges == nil")
	}

	successors := make(map[V]set.Set[V], vertices.Size())
	predecessors := make(map[V]set.Set[V], vertices.Size())

	for _, vertex := range vertices.Elements() {
		successors[vertex] = set.NewEmptyMapSet[V]()
		predecessors[vertex] = set.NewEmptyMapSet[V]()
	}

	for _, edge := range edges.Elements() {
		u := edge.Source()
		v := edge.Target()

		if u == v {
			err := fmt.Errorf(
				"loop edges are not allowed (%+v)",
				edge,
			)

			return nil, err
		}

		if !vertices.Contains(u) || !vertices.Contains(v) {
			err := fmt.Errorf(
				"source or target vertex of (%+v) is not present in vertices",
				edge,
			)

			return nil, err
		}

		successors[u].Add(v)
		predecessors[v].Add(u)
	}

	return &adjacencyListSimpleDigraph[V]{successors, predecessors}, nil
}

func (digraph *adjacencyListSimpleDigraph[V]) Vertices() set.Set[V] {
	vertices := make([]V, len(digraph.successors))

	i := 0
	for vertex := range digraph.successors {
		vertices[i] = vertex
		i++
	}

	return set.NewMapSetFromSlice(vertices)
}

func (digraph *adjacencyListSimpleDigraph[V]) Edges() set.Set[graph.Edge[V]] {
	edges := make([]graph.Edge[V], 0)

	for vertex, successors := range digraph.successors {
		for _, successor := range successors.Elements() {
			edges = append(edges, graph.NewEdge(vertex, successor))
		}
	}

	return set.NewMapSetFromSlice(edges)
}

func (digraph *adjacencyListSimpleDigraph[V]) Successors(vertex V) set.Set[V] {
	successors := set.NewEmptyMapSet[V]()

	if _, isPresent := digraph.successors[vertex]; !isPresent {
		return successors
	}

	for _, vertexSuccessor := range digraph.successors[vertex].Elements() {
		successors.Add(vertexSuccessor)
	}

	return successors
}

func (digraph *adjacencyListSimpleDigraph[V]) Predecessors(vertex V) set.Set[V] {
	predecessors := set.NewEmptyMapSet[V]()

	if _, isPresent := digraph.predecessors[vertex]; !isPresent {
		return predecessors
	}

	for _, vertexPredecessor := range digraph.predecessors[vertex].Elements() {
		predecessors.Add(vertexPredecessor)
	}

	return predecessors
}

func (digraph *adjacencyListSimpleDigraph[V]) Edge(
	source V,
	target V,
) *graph.Edge[V] {
	_, sourceIsPresent := digraph.successors[source]
	_, targetIsPresent := digraph.successors[target]

	if !sourceIsPresent || !targetIsPresent {
		return nil
	}

	for _, sourceSuccessor := range digraph.successors[source].Elements() {
		if sourceSuccessor == target {
			edge := graph.NewEdge(source, target)
			return &edge
		}
	}

	return nil
}

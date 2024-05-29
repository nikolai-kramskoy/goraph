package maxflow

import (
	"goraph/graph"
	"goraph/graph/digraph/simpledigraph"
)

// SimpleFlowNetwork https://en.wikipedia.org/wiki/Flow_network
type SimpleFlowNetwork[V graph.Vertex] struct {
	simpledigraph.SimpleDigraph[V]

	// S vertex must be present in graph.SimpleDigraph.
	S V

	// T vertex must be present in graph.SimpleDigraph.
	T V

	// Capacity must have a mapping for every edge in graph.SimpleDigraph.
	Capacity Capacity[V]

	// Flow must have a mapping for every edge in graph.SimpleDigraph.
	Flow Flow[V]
}

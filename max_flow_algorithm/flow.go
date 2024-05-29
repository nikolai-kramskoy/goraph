package max_flow_algorithm

import "goraph/graph"

// Flow maps graph.Edge to its integer non-negative uint32 flow.
//
// It has to satisfy given flow definition given in this article:
// https://en.wikipedia.org/wiki/Flow_network#Flows
type Flow[V graph.Vertex] map[graph.Edge[V]]uint32

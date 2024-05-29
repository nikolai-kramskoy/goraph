package max_flow_algorithm

import "goraph/graph"

// Capacity maps graph.Edge to its integer non-negative uint32 capacity.
type Capacity[V graph.Vertex] map[graph.Edge[V]]uint32

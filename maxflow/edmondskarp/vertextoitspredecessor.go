package edmondskarp

import "goraph/graph"

// vertexToPredecessor maps graph.Vertex to its predecessor (parent) graph.Vertex.
type vertexToPredecessor[V graph.Vertex] map[V]V

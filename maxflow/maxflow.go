package maxflow

import "goraph/graph"

// MaxFlow interface represents a max flow algorithm with single method
// that computes max Flow in SimpleFlowNetwork.
//
// No implementation can mutate SimpleFlowNetwork in any way.
//
// https://en.wikipedia.org/wiki/Maximum_flow_problem
type MaxFlow[V graph.Vertex] interface {
	// Compute computes max Flow in this SimpleFlowNetwork.
	//
	// If the graph is not st-connected, then max Flow returned
	// by this method == SimpleFlowNetwork.Flow.
	Compute(network *SimpleFlowNetwork[V]) (maxFlow Flow[V], err error)
}

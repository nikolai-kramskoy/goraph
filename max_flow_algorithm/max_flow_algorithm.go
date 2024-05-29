package max_flow_algorithm

import "goraph/graph"

// MaxFlowAlgorithm interface represents a max flow algorithm with single method
// that computes max Flow in SimpleFlowNetwork.
//
// No implementation can mutate SimpleFlowNetwork in any way.
//
// https://en.wikipedia.org/wiki/Maximum_flow_problem
type MaxFlowAlgorithm[V graph.Vertex] interface {
	// ComputeMaxFlow computes max Flow in this SimpleFlowNetwork.
	//
	// If the graph is not st-connected, then max Flow returned
	// by this method == SimpleFlowNetwork.Flow.
	ComputeMaxFlow(network *SimpleFlowNetwork[V]) (maxFlow Flow[V], err error)
}

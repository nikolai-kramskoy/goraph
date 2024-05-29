package edmondskarp

import (
	"goraph/graph"
	"goraph/graph/digraph/simpledigraph"
	al "goraph/graph/digraph/simpledigraph/adjacencylist"
	mf "goraph/maxflow"
	"math"

	"github.com/nikolai-kramskoy/go-data-structures/queue/slicequeue"
	"github.com/nikolai-kramskoy/go-data-structures/set"
	"github.com/nikolai-kramskoy/go-data-structures/set/mapset"
)

type edmondsKarp[V graph.Vertex] struct{}

var _ mf.MaxFlow[struct{}] = (*edmondsKarp[struct{}])(nil)

// NewEdmondsKarp creates an Edmonds-Karp algorithm
// implementation of max_flow_algorithm.MaxFlowAlgorithm.
//
// This implementation is immutable and thread-safe.
//
// https://en.wikipedia.org/wiki/Edmonds-Karp_algorithm
func NewEdmondsKarp[V graph.Vertex]() mf.MaxFlow[V] {
	return edmondsKarp[V]{}
}

func (algorithm edmondsKarp[V]) Compute(
	network *mf.SimpleFlowNetwork[V],
) (mf.Flow[V], error) {
	assertPreconditions(network)

	vertices := network.Vertices()
	currentFlow := copyFlow(network.Flow)

	// first iteration setup

	residualNetworkEdges, residualNetworkCapacity := residualSimpleFlowNetwork(network, currentFlow)
	residualDigraph, _ := al.NewAdjacencyListSimpleDigraph(vertices, residualNetworkEdges)
	vertexToPredecessor := breadthFirstSearch(residualDigraph, network.S, network.T)
	_, tPredecessorIsPresent := vertexToPredecessor[network.T]

	// augmenting (s,t)-path has been found in residual network
	for tPredecessorIsPresent {
		// find min edge capacity in residual network
		delta := computeDelta(network.T, vertexToPredecessor, residualNetworkCapacity)

		// now we need to increase currentFlow along augmenting (s,t)-semi-path in original network

		v := network.T
		u, uIsPresent := vertexToPredecessor[v]

		for uIsPresent {
			uv := network.Edge(u, v)
			vu := network.Edge(v, u)

			uvIsNotSaturated := uv != nil
			vuIsNotFlowless := vu != nil

			switch {
			case uvIsNotSaturated && vuIsNotFlowless:
				uvDelta := delta
				uvCurrentFlow := currentFlow[*uv]

				// if uv overflow did happen (iff vu != nil and currentFlow(vu) > 0)
				if uvCurrentFlow+delta > network.Capacity[*uv] {
					uvOverflow := uvCurrentFlow + delta - network.Capacity[*uv]
					uvDelta = delta - uvOverflow

					currentFlow[*vu] = currentFlow[*vu] - uvOverflow
				}

				currentFlow[*uv] = uvCurrentFlow + uvDelta

			case uvIsNotSaturated:
				currentFlow[*uv] = currentFlow[*uv] + delta

			case vuIsNotFlowless:
				currentFlow[*vu] = currentFlow[*vu] - delta
			}

			v = u
			u, uIsPresent = vertexToPredecessor[v]
		}

		// setup next iteration

		residualNetworkEdges, residualNetworkCapacity = residualSimpleFlowNetwork(network, currentFlow)
		residualDigraph, _ = al.NewAdjacencyListSimpleDigraph(vertices, residualNetworkEdges)
		vertexToPredecessor = breadthFirstSearch(residualDigraph, network.S, network.T)
		_, tPredecessorIsPresent = vertexToPredecessor[network.T]
	}

	return currentFlow, nil
}

func assertPreconditions[V graph.Vertex](network *mf.SimpleFlowNetwork[V]) {
	if network == nil {
		panic("network == nil")
	}
	if network.SimpleDigraph == nil {
		panic("network.SimpleDigraph == nil")
	}
	if network.Capacity == nil {
		panic("network.Capacity == nil")
	}
	if network.Flow == nil {
		panic("network.Flow == nil")
	}
}

func copyFlow[V graph.Vertex](flow mf.Flow[V]) mf.Flow[V] {
	copiedFlow := make(mf.Flow[V], len(flow))

	for edge, edgeFlow := range flow {
		copiedFlow[edge] = edgeFlow
	}

	return copiedFlow
}

// constructs residual network represented by pair (edges, capacity)
func residualSimpleFlowNetwork[V graph.Vertex](
	network *mf.SimpleFlowNetwork[V],
	currentFlow mf.Flow[V],
) (set.Set[graph.Edge[V]], mf.Capacity[V]) {
	edges := mapset.New[graph.Edge[V]]()
	capacity := make(mf.Capacity[V])

	// iterate over all permutations of vertices

	vertices := network.Vertices().Elements()
	verticesLen := len(vertices)

	for i := 0; i < verticesLen; i++ {
		for j := 0; j < verticesLen; j++ {
			// In simple flow network we don't care about loops
			if i == j {
				continue
			}

			u := vertices[i]
			v := vertices[j]

			uv := network.Edge(u, v)
			vu := network.Edge(v, u)

			uvIsNotSaturated := uv != nil && currentFlow[*uv] < network.Capacity[*uv]
			vuIsNotFlowless := vu != nil && currentFlow[*vu] > 0

			if uvIsNotSaturated || vuIsNotFlowless {
				residualUv := graph.NewEdge(u, v)

				edges.Add(residualUv)

				switch {
				case uvIsNotSaturated && vuIsNotFlowless:
					capacity[residualUv] = network.Capacity[*uv] - currentFlow[*uv] + currentFlow[*vu]

				case uvIsNotSaturated:
					capacity[residualUv] = network.Capacity[*uv] - currentFlow[*uv]

				case vuIsNotFlowless:
					capacity[residualUv] = currentFlow[*vu]
				}
			}
		}
	}

	return edges, capacity
}

func breadthFirstSearch[V graph.Vertex](
	digraph simpledigraph.SimpleDigraph[V],
	s V,
	t V,
) vertexToPredecessor[V] {
	vertexQueue := slicequeue.New[V]()
	visitedVertices := mapset.New[V]()
	vertexToPredecessor := make(vertexToPredecessor[V])

	vertexQueue.Push(s)
	visitedVertices.Add(s)

	for !vertexQueue.IsEmpty() {
		u := vertexQueue.Pop()

		for _, v := range digraph.Successors(u).Elements() {
			if !visitedVertices.Contains(v) {
				vertexToPredecessor[v] = u

				if v == t {
					return vertexToPredecessor
				}

				vertexQueue.Push(v)
				visitedVertices.Add(v)
			}
		}
	}

	return vertexToPredecessor
}

func computeDelta[V graph.Vertex](
	t V,
	vertexToPredecessor vertexToPredecessor[V],
	residualNetworkCapacity mf.Capacity[V],
) uint32 {
	v := t
	u, uIsPresent := vertexToPredecessor[v]
	var delta uint32 = math.MaxUint32

	for uIsPresent {
		residualNetworkUv := graph.NewEdge(u, v)
		residualNetworkUvCapacity := residualNetworkCapacity[residualNetworkUv]

		if residualNetworkUvCapacity < delta {
			delta = residualNetworkUvCapacity
		}

		v = u
		u, uIsPresent = vertexToPredecessor[v]
	}

	return delta
}

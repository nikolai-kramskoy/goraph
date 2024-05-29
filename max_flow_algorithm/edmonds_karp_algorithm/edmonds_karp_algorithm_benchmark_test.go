package edmonds_karp_algorithm

import (
	"fmt"
	"goraph/collection/set"
	"goraph/graph"
	ald "goraph/graph/adjacency_list_digraph"
	mfa "goraph/max_flow_algorithm"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func BenchmarkEdmondsKarpAlgorithm_ComputeMaxFlow(b *testing.B) {
	rand.Seed(time.Now().Unix())

	amountOfVertices, amountOfEdges, graphIsComplete := getParameters()

	fmt.Printf(
		"amountOfVertices = %d, amountOfEdges = %d, graphIsComplete = %t\n",
		amountOfVertices,
		amountOfEdges,
		graphIsComplete,
	)

	simpleFlowNetwork := generateRandomSimpleFlowNetwork(amountOfVertices, amountOfEdges, graphIsComplete)

	maxFlowAlgorithm := NewEdmondsKarpAlgorithm[int]()

	b.ResetTimer()

	maxFlow, maxFlowAlgorithmErr := maxFlowAlgorithm.ComputeMaxFlow(simpleFlowNetwork)

	b.StopTimer()

	if maxFlowAlgorithmErr != nil {
		panic(maxFlowAlgorithmErr)
	}

	maxFlowValue := computeMaxFlowValue(simpleFlowNetwork, maxFlow)

	fmt.Printf("maxFlowValue = %d\n", maxFlowValue)
}

func computeMaxFlowValue(network *mfa.SimpleFlowNetwork[int], maxFlow mfa.Flow[int]) int64 {
	sSuccessors := network.Successors(network.S)
	sPredecessors := network.Predecessors(network.S)

	var maxFlowValue int64 = 0

	u := network.S

	for _, v := range sSuccessors.Elements() {
		maxFlowValue += int64(maxFlow[*network.Edge(u, v)])
	}

	for _, v := range sPredecessors.Elements() {
		maxFlowValue -= int64(maxFlow[*network.Edge(v, u)])
	}

	return maxFlowValue
}

func getParameters() (int, int, bool) {
	argsLastIndex := len(os.Args) - 1

	amountOfVertices, err1 := strconv.Atoi(os.Args[argsLastIndex-2])

	if err1 != nil {
		panic(err1)
	}

	amountOfEdges, err2 := strconv.Atoi(os.Args[argsLastIndex-1])

	if err2 != nil {
		panic(err2)
	}

	graphIsComplete, err3 := strconv.ParseBool(os.Args[argsLastIndex])

	if err3 != nil {
		panic(err3)
	}

	return amountOfVertices, amountOfEdges, graphIsComplete
}

// Vertices are labeled from 1 to amountOfVertices.
func generateRandomSimpleFlowNetwork(
	amountOfVertices int,
	amountOfEdges int,
	graphIsComplete bool,
) *mfa.SimpleFlowNetwork[int] {
	vertices := newVertices(amountOfVertices)

	var edges set.Set[graph.Edge[int]]
	var capacity mfa.Capacity[int]
	var flow mfa.Flow[int]

	if graphIsComplete {
		edges, capacity, flow = newEdgesCapacityFlowCompleteGraph(amountOfVertices)
	} else {
		edges, capacity, flow = newEdgesCapacityFlow(amountOfVertices, amountOfEdges)
	}

	simpleDigraph, newSimpleDigraphErr := ald.NewAdjacencyListSimpleDigraph(vertices, edges)

	if newSimpleDigraphErr != nil {
		panic(newSimpleDigraphErr)
	}

	s := rand.Intn(amountOfVertices) + 1
	t := rand.Intn(amountOfVertices) + 1

	for s == t {
		t = rand.Intn(amountOfVertices) + 1
	}

	return &mfa.SimpleFlowNetwork[int]{
		SimpleDigraph: simpleDigraph,
		S:             s,
		T:             t,
		Capacity:      capacity,
		Flow:          flow,
	}
}

func newVertices(amountOfVertices int) set.Set[int] {
	vertices := set.NewEmptyMapSet[int]()

	for i := 1; i <= amountOfVertices; i++ {
		vertices.Add(i)
	}

	return vertices
}

func newEdgesCapacityFlowCompleteGraph(amountOfVertices int) (
	set.Set[graph.Edge[int]],
	mfa.Capacity[int],
	mfa.Flow[int],
) {
	amountOfEdges := amountOfVertices * (amountOfVertices - 1)

	edgesSliceIndex := 0
	edgesSlice := make([]graph.Edge[int], amountOfEdges)
	capacity := make(mfa.Capacity[int], amountOfEdges)
	flow := make(mfa.Flow[int], amountOfEdges)

	for u := 1; u <= amountOfVertices; u++ {
		for v := 1; v <= amountOfVertices; v++ {
			if u == v {
				continue
			}

			uv := graph.NewEdge(u, v)
			uvCapacity, uvFlow := randCapacityAndFlowValues()

			edgesSlice[edgesSliceIndex] = uv
			edgesSliceIndex++
			capacity[uv] = uvCapacity
			flow[uv] = uvFlow
		}
	}

	return set.NewMapSetFromSlice(edgesSlice), capacity, flow
}

func newEdgesCapacityFlow(
	amountOfVertices int,
	amountOfEdges int,
) (set.Set[graph.Edge[int]], mfa.Capacity[int], mfa.Flow[int]) {
	edgesSliceIndex := 0
	edgesSlice := make([]graph.Edge[int], amountOfEdges)
	capacity := make(mfa.Capacity[int], amountOfEdges)
	flow := make(mfa.Flow[int], amountOfEdges)

	for i := 1; i <= amountOfEdges; i++ {
		u := rand.Intn(amountOfVertices) + 1
		v := rand.Intn(amountOfVertices) + 1

		for u == v {
			v = rand.Intn(amountOfVertices) + 1
		}

		uv := graph.NewEdge(u, v)
		uvCapacity, uvFlow := randCapacityAndFlowValues()

		edgesSlice[edgesSliceIndex] = uv
		edgesSliceIndex++
		capacity[uv] = uvCapacity
		flow[uv] = uvFlow
	}

	return set.NewMapSetFromSlice(edgesSlice), capacity, flow
}

func randCapacityAndFlowValues() (uint32, uint32) {
	// to avoid overflow in (capacityValue + 1) expression
	capacityValue := uint32(rand.Int31())

	var flowValue uint32

	if capacityValue == 0 {
		flowValue = 0
	} else {
		flowValue = rand.Uint32() % (capacityValue + 1)
	}

	return capacityValue, flowValue
}

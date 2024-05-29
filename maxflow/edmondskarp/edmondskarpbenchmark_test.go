package edmondskarp

import (
	"fmt"
	"goraph/graph"
	al "goraph/graph/digraph/simpledigraph/adjacencylist"
	mf "goraph/maxflow"
	"math/rand"
	"testing"

	"github.com/nikolai-kramskoy/go-data-structures/set"
	"github.com/nikolai-kramskoy/go-data-structures/set/mapset"
)

func BenchmarkEdmondsKarp_Compute_1(b *testing.B) {
	amountOfVertices := 10
	amountOfEdges := 50

	fmt.Printf(
		"amountOfVertices = %d, amountOfEdges = %d, graphIsComplete = %t ",
		amountOfVertices,
		amountOfEdges,
		false,
	)

	edmondsKarp_Compute_Benchmark(b, generateSimpleFlowNetwork(amountOfVertices, amountOfEdges))
}

func BenchmarkEdmondsKarp_Compute_2(b *testing.B) {
	amountOfVertices := 100
	amountOfEdges := 5000

	fmt.Printf(
		"amountOfVertices = %d, amountOfEdges = %d, graphIsComplete = %t ",
		amountOfVertices,
		amountOfEdges,
		false,
	)

	edmondsKarp_Compute_Benchmark(b, generateSimpleFlowNetwork(amountOfVertices, amountOfEdges))
}

func BenchmarkEdmondsKarp_Compute_3(b *testing.B) {
	amountOfVertices := 250
	amountOfEdges := amountOfVertices * amountOfVertices

	fmt.Printf(
		"amountOfVertices = %d, amountOfEdges = %d, graphIsComplete = %t ",
		amountOfVertices,
		amountOfEdges,
		true,
	)

	edmondsKarp_Compute_Benchmark(b, generateCompleteSimpleFlowNetwork(amountOfVertices))
}

func edmondsKarp_Compute_Benchmark(b *testing.B, network *mf.SimpleFlowNetwork[int]) {
	edmondsKarp := NewEdmondsKarp[int]()

	b.ResetTimer()

	maxFlow, err := edmondsKarp.Compute(network)

	b.StopTimer()

	if err != nil {
		panic(err)
	}

	maxFlowValue := computeMaxFlowValue(network, maxFlow)

	fmt.Printf("maxFlowValue = %d\n", maxFlowValue)
}

func computeMaxFlowValue(network *mf.SimpleFlowNetwork[int], maxFlow mf.Flow[int]) int64 {
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

// Vertices are labeled from 1 to amountOfVertices.
func generateSimpleFlowNetwork(
	amountOfVertices int,
	amountOfEdges int,
) *mf.SimpleFlowNetwork[int] {
	edges, capacity, flow := newEdgesCapacityFlow(amountOfVertices, amountOfEdges)

	return generateSimpleFlowNetworkHelper(amountOfVertices, edges, capacity, flow)
}

// Vertices are labeled from 1 to amountOfVertices.
func generateCompleteSimpleFlowNetwork(
	amountOfVertices int,
) *mf.SimpleFlowNetwork[int] {
	edges, capacity, flow := newEdgesCapacityFlowCompleteGraph(amountOfVertices)

	return generateSimpleFlowNetworkHelper(amountOfVertices, edges, capacity, flow)
}

func generateSimpleFlowNetworkHelper(
	amountOfVertices int,
	edges set.Set[graph.Edge[int]],
	capacity mf.Capacity[int],
	flow mf.Flow[int],
) *mf.SimpleFlowNetwork[int] {
	vertices := newVertices(amountOfVertices)

	simpleDigraph, err := al.NewAdjacencyListSimpleDigraph(vertices, edges)

	if err != nil {
		panic(err)
	}

	s := rand.Intn(amountOfVertices) + 1
	t := rand.Intn(amountOfVertices) + 1

	for s == t {
		t = rand.Intn(amountOfVertices) + 1
	}

	return &mf.SimpleFlowNetwork[int]{
		SimpleDigraph: simpleDigraph,
		S:             s,
		T:             t,
		Capacity:      capacity,
		Flow:          flow,
	}
}

func newVertices(amountOfVertices int) set.Set[int] {
	vertices := mapset.New[int]()

	for i := 1; i <= amountOfVertices; i++ {
		vertices.Add(i)
	}

	return vertices
}

func newEdgesCapacityFlowCompleteGraph(amountOfVertices int) (
	set.Set[graph.Edge[int]],
	mf.Capacity[int],
	mf.Flow[int],
) {
	amountOfEdges := amountOfVertices * (amountOfVertices - 1)

	edgesSliceIndex := 0
	edgesSlice := make([]graph.Edge[int], amountOfEdges)
	capacity := make(mf.Capacity[int], amountOfEdges)
	flow := make(mf.Flow[int], amountOfEdges)

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

	return mapset.NewFromElements(edgesSlice...), capacity, flow
}

func newEdgesCapacityFlow(
	amountOfVertices int,
	amountOfEdges int,
) (set.Set[graph.Edge[int]], mf.Capacity[int], mf.Flow[int]) {
	edgesSliceIndex := 0
	edgesSlice := make([]graph.Edge[int], amountOfEdges)
	capacity := make(mf.Capacity[int], amountOfEdges)
	flow := make(mf.Flow[int], amountOfEdges)

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

	return mapset.NewFromElements(edgesSlice...), capacity, flow
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

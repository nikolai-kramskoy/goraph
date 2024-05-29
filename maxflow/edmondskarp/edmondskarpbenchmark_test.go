package edmondskarp

import (
	"fmt"
	"goraph/graph"
	al "goraph/graph/digraph/simpledigraph/adjacencylist"
	mf "goraph/maxflow"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/nikolai-kramskoy/go-data-structures/set"
	"github.com/nikolai-kramskoy/go-data-structures/set/mapset"
)

func BenchmarkEdmondsKarp_Compute(b *testing.B) {
	amountOfVertices, amountOfEdges, graphIsComplete := getParameters()

	fmt.Printf(
		"amountOfVertices = %d, amountOfEdges = %d, graphIsComplete = %t\n",
		amountOfVertices,
		amountOfEdges,
		graphIsComplete,
	)

	network := generateRandomSimpleFlowNetwork(amountOfVertices, amountOfEdges, graphIsComplete)

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
) *mf.SimpleFlowNetwork[int] {
	vertices := newVertices(amountOfVertices)

	var edges set.Set[graph.Edge[int]]
	var capacity mf.Capacity[int]
	var flow mf.Flow[int]

	if graphIsComplete {
		edges, capacity, flow = newEdgesCapacityFlowCompleteGraph(amountOfVertices)
	} else {
		edges, capacity, flow = newEdgesCapacityFlow(amountOfVertices, amountOfEdges)
	}

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

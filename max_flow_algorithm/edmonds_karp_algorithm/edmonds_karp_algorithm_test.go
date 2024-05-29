package edmonds_karp_algorithm

import (
	"github.com/stretchr/testify/assert"
	"goraph/collection/set"
	"goraph/graph"
	ald "goraph/graph/adjacency_list_digraph"
	mfa "goraph/max_flow_algorithm"
	"testing"
)

// I"ve used example from this website
// https://en.wikipedia.org/wiki/Edmonds-Karp_algorithm#Example
func newExampleSimpleFlowNetwork() (*mfa.SimpleFlowNetwork[string], mfa.Flow[string]) {
	a, b, c, d, e, f, g := "A", "B", "C", "D", "E", "F", "G"

	vertices := set.NewMapSet(a, b, c, d, e, f, g)

	ab, ad := graph.NewEdge(a, b), graph.NewEdge(a, d)
	bc := graph.NewEdge(b, c)
	ca, cd, ce := graph.NewEdge(c, a), graph.NewEdge(c, d), graph.NewEdge(c, e)
	de, df := graph.NewEdge(d, e), graph.NewEdge(d, f)
	eb, eg := graph.NewEdge(e, b), graph.NewEdge(e, g)
	fg := graph.NewEdge(f, g)

	edges := set.NewMapSet(
		ab, ad,
		bc,
		ca, cd, ce,
		de, df,
		eb, eg,
		fg,
	)

	simpleDigraph, _ := ald.NewAdjacencyListSimpleDigraph(vertices, edges)

	capacity := mfa.Capacity[string]{
		ab: 3, ad: 3,
		bc: 4,
		ca: 3, cd: 1, ce: 2,
		de: 2, df: 6,
		eb: 1, eg: 1,
		fg: 9,
	}

	flow := mfa.Flow[string]{
		ab: 0, ad: 0,
		bc: 0,
		ca: 0, cd: 0, ce: 0,
		de: 0, df: 0,
		eb: 0, eg: 0,
		fg: 0,
	}

	expectedMaxFlow := mfa.Flow[string]{
		ab: 2, ad: 3,
		bc: 2,
		ca: 0, cd: 1, ce: 1,
		de: 0, df: 4,
		eb: 0, eg: 1,
		fg: 4,
	}

	return &mfa.SimpleFlowNetwork[string]{
			SimpleDigraph: simpleDigraph,
			S:             a,
			T:             g,
			Capacity:      capacity,
			Flow:          flow,
		},
		expectedMaxFlow
}

func TestEdmondsKarpAlgorithm_Compute(t *testing.T) {
	simpleFlowNetwork, expectedMaxFlow := newExampleSimpleFlowNetwork()

	maxFlowAlgorithm := NewEdmondsKarpAlgorithm[string]()

	maxFlow, err := maxFlowAlgorithm.ComputeMaxFlow(simpleFlowNetwork)

	assert.NotNil(t, maxFlow)
	assert.Nil(t, err)

	assert.Equal(t, expectedMaxFlow, maxFlow)
}

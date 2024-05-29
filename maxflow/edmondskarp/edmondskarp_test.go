package edmondskarp

import (
	"goraph/graph"
	al "goraph/graph/digraph/simpledigraph/adjacencylist"
	mf "goraph/maxflow"
	"testing"

	"github.com/nikolai-kramskoy/go-data-structures/set/mapset"

	"github.com/stretchr/testify/assert"
)

// I"ve used example from this website
// https://en.wikipedia.org/wiki/Edmonds-Karp_algorithm#Example
func newExampleSimpleFlowNetwork() (*mf.SimpleFlowNetwork[string], mf.Flow[string]) {
	a, b, c, d, e, f, g := "A", "B", "C", "D", "E", "F", "G"

	vertices := mapset.NewFromElements(a, b, c, d, e, f, g)

	ab, ad := graph.NewEdge(a, b), graph.NewEdge(a, d)
	bc := graph.NewEdge(b, c)
	ca, cd, ce := graph.NewEdge(c, a), graph.NewEdge(c, d), graph.NewEdge(c, e)
	de, df := graph.NewEdge(d, e), graph.NewEdge(d, f)
	eb, eg := graph.NewEdge(e, b), graph.NewEdge(e, g)
	fg := graph.NewEdge(f, g)

	edges := mapset.NewFromElements(
		ab, ad,
		bc,
		ca, cd, ce,
		de, df,
		eb, eg,
		fg,
	)

	simpleDigraph, _ := al.NewAdjacencyListSimpleDigraph(vertices, edges)

	capacity := mf.Capacity[string]{
		ab: 3, ad: 3,
		bc: 4,
		ca: 3, cd: 1, ce: 2,
		de: 2, df: 6,
		eb: 1, eg: 1,
		fg: 9,
	}

	flow := mf.Flow[string]{
		ab: 0, ad: 0,
		bc: 0,
		ca: 0, cd: 0, ce: 0,
		de: 0, df: 0,
		eb: 0, eg: 0,
		fg: 0,
	}

	expectedMaxFlow := mf.Flow[string]{
		ab: 2, ad: 3,
		bc: 2,
		ca: 0, cd: 1, ce: 1,
		de: 0, df: 4,
		eb: 0, eg: 1,
		fg: 4,
	}

	return &mf.SimpleFlowNetwork[string]{
			SimpleDigraph: simpleDigraph,
			S:             a,
			T:             g,
			Capacity:      capacity,
			Flow:          flow,
		},
		expectedMaxFlow
}

func TestEdmondsKarp_Compute(t *testing.T) {
	network, expectedMaxFlow := newExampleSimpleFlowNetwork()

	edmondsKarp := NewEdmondsKarp[string]()

	actualMaxFlow, err := edmondsKarp.Compute(network)

	assert.NotNil(t, actualMaxFlow)
	assert.Nil(t, err)

	assert.Equal(t, expectedMaxFlow, actualMaxFlow)
}

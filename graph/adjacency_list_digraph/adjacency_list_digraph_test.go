package adjacency_list_digraph

import (
	"github.com/stretchr/testify/assert"
	"goraph/collection/set"
	"goraph/graph"
	"testing"
)

var vertices = set.NewMapSet(1, 2, 3, 4)
var edges = set.NewMapSet(
	graph.NewEdge(1, 2),
	graph.NewEdge(1, 4),
	graph.NewEdge(2, 1),
	graph.NewEdge(2, 3),
	graph.NewEdge(2, 4),
	graph.NewEdge(3, 4),
)

func TestAdjacencyListDigraph(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	digraphVertices := simpleDigraph.Vertices()
	assert.Equal(t, vertices, digraphVertices)

	digraphEdges := simpleDigraph.Edges()
	assert.Equal(t, edges, digraphEdges)

	// (2, 3) edge exists

	twoThreeEdge := simpleDigraph.Edge(2, 3)
	assert.NotNil(t, twoThreeEdge)
	assert.Equal(t, graph.NewEdge(2, 3), *twoThreeEdge)

	// (1, 1) loop edge does not exist

	oneOneEdge := simpleDigraph.Edge(1, 1)
	assert.Nil(t, oneOneEdge)
}

func TestAdjacencyListDigraph_Successors(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	twoSuccessors := simpleDigraph.Successors(2)

	assert.Equal(t, set.NewMapSet(1, 3, 4), twoSuccessors)
}

func TestAdjacencyListDigraph_Successors2(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	twoSuccessors := simpleDigraph.Successors(-2)

	assert.Equal(t, set.NewEmptyMapSet[int](), twoSuccessors)
}

func TestAdjacencyListDigraph_Predecessors(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	twoPredecessors := simpleDigraph.Predecessors(2)

	assert.Equal(t, set.NewMapSet(1), twoPredecessors)
}

func TestAdjacencyListDigraph_Predecessor2(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	twoPredecessors := simpleDigraph.Predecessors(-6)

	assert.Equal(t, set.NewEmptyMapSet[int](), twoPredecessors)
}

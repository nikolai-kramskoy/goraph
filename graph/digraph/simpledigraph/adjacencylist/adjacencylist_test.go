package adjacencylist

import (
	"goraph/graph"
	"testing"

	"github.com/nikolai-kramskoy/go-data-structures/set/mapset"

	"github.com/stretchr/testify/assert"
)

var vertices = mapset.NewFromElements(1, 2, 3, 4)
var edges = mapset.NewFromElements(
	graph.NewEdge(1, 2),
	graph.NewEdge(1, 4),
	graph.NewEdge(2, 1),
	graph.NewEdge(2, 3),
	graph.NewEdge(2, 4),
	graph.NewEdge(3, 4),
)

func TestAdjacencyListSimpleDigraph(t *testing.T) {
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

func TestAdjacencyListSimpleDigraph_Successors(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	twoSuccessors := simpleDigraph.Successors(2)

	assert.Equal(t, mapset.NewFromElements(1, 3, 4), twoSuccessors)
}

func TestAdjacencyListSimpleDigraph_Successors2(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	twoSuccessors := simpleDigraph.Successors(-2)

	assert.Equal(t, mapset.New[int](), twoSuccessors)
}

func TestAdjacencyListSimpleDigraph_Predecessors(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	twoPredecessors := simpleDigraph.Predecessors(2)

	assert.Equal(t, mapset.NewFromElements(1), twoPredecessors)
}

func TestAdjacencyListSimpleDigraph_Predecessor2(t *testing.T) {
	simpleDigraph, err := NewAdjacencyListSimpleDigraph(vertices, edges)
	assert.NotNil(t, simpleDigraph)
	assert.Nil(t, err)

	twoPredecessors := simpleDigraph.Predecessors(-6)

	assert.Equal(t, mapset.New[int](), twoPredecessors)
}

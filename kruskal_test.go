package minimumspanningtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EasyKruskal(t *testing.T) {
	edges := []Edge{
		{Source: 0, Destination: 1, Weight: 1},
		{Source: 0, Destination: 2, Weight: 3},
		{Source: 0, Destination: 3, Weight: 4},
		{Source: 1, Destination: 2, Weight: 2},
		{Source: 1, Destination: 3, Weight: 5},
		{Source: 2, Destination: 3, Weight: 6},
	}

	expectPath := []Edge{
		{Source: 0, Destination: 1, Weight: 1},
		{Source: 1, Destination: 2, Weight: 2},
		{Source: 0, Destination: 3, Weight: 4},
	}

	kruskal := NewKruskal(edges)

	result := kruskal.GetMinimumSpanningTree()

	require.Equal(t, expectPath, result)
}

func Test_ComplextKruskal(t *testing.T) {
	edges := []Edge{
		{Source: 0, Destination: 1, Weight: 4},
		{Source: 0, Destination: 7, Weight: 8},
		{Source: 1, Destination: 2, Weight: 8},
		{Source: 1, Destination: 7, Weight: 11},
		{Source: 2, Destination: 3, Weight: 7},
		{Source: 2, Destination: 8, Weight: 2},
		{Source: 2, Destination: 5, Weight: 4},
		{Source: 3, Destination: 4, Weight: 9},
		{Source: 3, Destination: 5, Weight: 14},
		{Source: 4, Destination: 5, Weight: 10},
		{Source: 5, Destination: 6, Weight: 2},
		{Source: 6, Destination: 7, Weight: 1},
		{Source: 6, Destination: 8, Weight: 6},
		{Source: 7, Destination: 8, Weight: 7},
	}

	expectPath := []Edge{
		{Source: 6, Destination: 7, Weight: 1},
		{Source: 2, Destination: 8, Weight: 2},
		{Source: 5, Destination: 6, Weight: 2},
		{Source: 2, Destination: 5, Weight: 4},
		{Source: 0, Destination: 1, Weight: 4},
		{Source: 2, Destination: 3, Weight: 7},
		{Source: 0, Destination: 7, Weight: 8},
		{Source: 3, Destination: 4, Weight: 9},
	}

	kruskal := NewKruskal(edges)

	result := kruskal.GetMinimumSpanningTree()

	require.Equal(t, expectPath, result)
}

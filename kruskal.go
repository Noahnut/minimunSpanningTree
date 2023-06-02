package minimumspanningtree

import "sort"

type Edge struct {
	Source      int
	Destination int
	Weight      int
}

type Kruskal struct {
	edges []Edge
}

func NewKruskal(edges []Edge) *Kruskal {
	return &Kruskal{
		edges: edges,
	}
}

func (k *Kruskal) GetMinimumSpanningTree() []Edge {
	minSpannerTree := make([]Edge, 0)

	sort.Slice(k.edges, func(i, j int) bool {
		return k.edges[i].Weight < k.edges[j].Weight
	})

	unionSet := NewUnionSet(len(k.edges))

	for i := 0; i < len(k.edges); i++ {
		if unionSet.Find(k.edges[i].Source) != unionSet.Find(k.edges[i].Destination) {
			minSpannerTree = append(minSpannerTree, k.edges[i])
			unionSet.Union(k.edges[i].Source, k.edges[i].Destination)
		}
	}

	return minSpannerTree
}

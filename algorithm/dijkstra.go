package algorithm

import (
	"sort"

	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid/node"
)

func getAllNodes(grid [][]*node.Node) []*node.Node {
	var nodes []*node.Node
	for _, r := range grid {
		nodes = append(nodes, r...)
	}
	return nodes
}

func updateUnvisitedNeighborNodes(n *node.Node, g [][]*node.Node) []*node.Node {
	unvisitedNeighbors := []*node.Node{}
	l := grid.Length / grid.BlockSize
	w := grid.Width / grid.BlockSize

	if n.Row > 0 {
		unvisitedNeighbors = append(unvisitedNeighbors, g[n.Row-1][n.Col])
	}
	if n.Row < l-1 {
		unvisitedNeighbors = append(unvisitedNeighbors, g[n.Row+1][n.Col])
	}
	if n.Col > 0 {
		unvisitedNeighbors = append(unvisitedNeighbors, g[n.Row][n.Col-1])
	}
	if n.Col < w-1 {
		unvisitedNeighbors = append(unvisitedNeighbors, g[n.Row][n.Col+1])
	}

	uNodes := []*node.Node{}

	for _, un := range unvisitedNeighbors {
		if !un.IsVisited {
			un.Distance = n.Distance + 1
			un.PreviousVisited = n
			uNodes = append(uNodes, un)
		}
	}

	return uNodes

}

func Dijkstra(v *grid.Visualiser) []*node.Node {
	var visitedNodes []*node.Node

	v.StartNode.Distance = 0
	nodes := getAllNodes(v.Grid)

	for len(nodes) != 0 {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Distance < nodes[j].Distance
		})
		n := nodes[0]
		nodes = nodes[1:]

		if n.IsBarrier {
			continue
		}

		if n.Distance == grid.Infinity {
			return visitedNodes
		}

		n.IsVisited = true
		visitedNodes = append(visitedNodes, n)

		if n == v.EndNode {
			return visitedNodes
		}

		updateUnvisitedNeighborNodes(n, v.Grid)
	}

	return visitedNodes
}

func GetNodesInShortPath(v *grid.Visualiser) []*node.Node {
	shortPathNodes := []*node.Node{}
	start := v.EndNode
	for start != nil {
		shortPathNodes = append(shortPathNodes, &node.Node{})
		copy(shortPathNodes[1:], shortPathNodes)
		shortPathNodes[0] = start
		start = start.PreviousVisited
	}

	return shortPathNodes
}

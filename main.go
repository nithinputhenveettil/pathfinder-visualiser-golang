package main

import (
	"fmt"
	"math"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type node struct {
	row             int32
	col             int32
	isStart         bool
	isFinish        bool
	isVisited       bool
	isAnimateOn     bool
	isBarrier       bool
	previousVisited *node
	distance        float64
}

var (
	width  int32 = 1300
	length int32 = 650
	FPS    int32 = 120

	BlockSize int32 = 25

	Infinity float64 = math.Inf(1)
)

var (
	start_row int32 = 10
	start_col int32 = 10

	end_row int32 = 10
	end_col int32 = 40
)

var (
	BG_COL = rl.White
	OT_COL = rl.Black

	ST_COL = rl.Red
)

var grid [][]*node
var startVisualise bool
var startNode, endNode *node
var vIndex int
var visitedNodes []*node

func getInitGrid() [][]*node {
	var g [][]*node
	var i, j int32

	l := length / BlockSize
	w := width / BlockSize

	fmt.Println(l, w)

	for i = 0; i < l; i++ {
		r := []*node{}
		for j = 0; j < w; j++ {
			n := &node{
				row:      i,
				col:      j,
				isStart:  j == start_col && i == start_row,
				isFinish: j == end_col && i == end_row,
				distance: Infinity,
			}
			r = append(r, n)
			if n.isStart {
				startNode = n
			}
			if n.isFinish {
				endNode = n
			}
		}
		g = append(g, r)
	}

	return g

}

func drawGrid(g [][]*node) {
	for _, r := range g {
		for _, n := range r {
			x := n.col * BlockSize
			y := n.row * BlockSize
			drawNodeOutline(x, y)
			if n.isStart {
				drawStartNode(x, y)
			}
			if n.isFinish {
				drawEndNode(x, y)
			}
			if n.isBarrier && !n.isStart && !n.isFinish {
				drawBarrierNode(x, y)
			}
			if n.isVisited && n.isAnimateOn && !n.isStart && !n.isFinish {
				drawVisitedNode(x, y)
			}
		}
	}
}

func animateVisited(nodes []*node) {
	if vIndex == len(nodes) {
		return
	}
	n := nodes[vIndex]
	n.isAnimateOn = true
	vIndex += 1
}

func drawNodeOutline(x, y int32) {
	rl.DrawRectangleLines(x+1, y+1, BlockSize-3, BlockSize-3, OT_COL)
}

func drawStartNode(x, y int32) {
	rl.DrawLineEx(rl.Vector2{X: float32(x + (BlockSize / 4)), Y: float32(y + (BlockSize / 4))}, rl.Vector2{X: float32(x + (BlockSize * 3 / 4)), Y: float32(y + (BlockSize / 2))}, 4.0, ST_COL)
	rl.DrawLineEx(rl.Vector2{X: float32(x + (BlockSize / 4)), Y: float32(y + (BlockSize * 3 / 4))}, rl.Vector2{X: float32(x + (BlockSize * 3 / 4)), Y: float32(y + (BlockSize / 2))}, 4.0, ST_COL)
}

func drawEndNode(x, y int32) {
	rl.DrawCircleV(rl.Vector2{X: float32(x + (BlockSize / 2)), Y: float32(y + (BlockSize / 2))}, 3, ST_COL)
	rl.DrawRing(rl.Vector2{X: float32(x + (BlockSize / 2)), Y: float32(y + (BlockSize / 2))}, 6, 10, 0, 360, 0, ST_COL)
}

func drawBarrierNode(x, y int32) {
	rl.DrawRectangle(x+1, y+1, BlockSize-3, BlockSize-3, OT_COL)
}

func drawVisitedNode(x, y int32) {
	rl.DrawRectangle(x+1, y+1, BlockSize-3, BlockSize-3, rl.SkyBlue)
}

func litsenMouseClick() {
	if rl.IsMouseButtonPressed(0) {
		points := rl.GetMousePosition()
		fmt.Println(points)
		x := (int32)(points.X / float32(BlockSize))
		y := (int32)(points.Y / float32(BlockSize))
		grid[y][x].isBarrier = !grid[y][x].isBarrier
		fmt.Println(points, x, y)
	}
}

func litsenKeyboardEvents() {
	if rl.IsKeyPressed(83) {
		startVisualise = true
	}
	if rl.IsKeyPressed(82) {
		// reset
		startVisualise = false
		vIndex = 0
		grid = getInitGrid()
		visitedNodes = []*node{}
	}
}

func getAllNodes(grid [][]*node) []*node {
	var nodes []*node
	for _, r := range grid {
		nodes = append(nodes, r...)
	}
	return nodes
}

func updateUnvisitedNeighborNodes(n *node, g [][]*node) []*node {
	unvisitedNeighbors := []*node{}
	l := length / BlockSize
	w := width / BlockSize

	if n.row > 0 {
		unvisitedNeighbors = append(unvisitedNeighbors, g[n.row-1][n.col])
	}
	if n.row < l-1 {
		unvisitedNeighbors = append(unvisitedNeighbors, g[n.row+1][n.col])
	}
	if n.col > 0 {
		unvisitedNeighbors = append(unvisitedNeighbors, g[n.row][n.col-1])
	}
	if n.col < w-1 {
		unvisitedNeighbors = append(unvisitedNeighbors, g[n.row][n.col+1])
	}

	uNodes := []*node{}

	for _, un := range unvisitedNeighbors {
		if !un.isVisited {
			un.distance = n.distance + 1
			uNodes = append(uNodes, un)
		}
	}

	return uNodes

}

func dijkstra(startNode, endNode *node, g [][]*node) []*node {
	var visitedNodes []*node

	startNode.distance = 0
	nodes := getAllNodes(g)

	for len(nodes) != 0 {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].distance < nodes[j].distance
		})
		n := nodes[0]
		nodes = nodes[1:]

		if n.isBarrier {
			continue
		}

		if n.distance == Infinity {
			return visitedNodes
		}

		n.isVisited = true
		visitedNodes = append(visitedNodes, n)

		if n == endNode {
			return visitedNodes
		}

		updateUnvisitedNeighborNodes(n, g)
	}

	return visitedNodes
}

func main() {
	grid = getInitGrid()
	rl.InitWindow(width, length, "Path Finder Visualiser")
	rl.SetTargetFPS(FPS)
	for !rl.WindowShouldClose() {
		litsenMouseClick()
		litsenKeyboardEvents()
		if startVisualise {
			visitedNodes = dijkstra(startNode, endNode, grid)
			fmt.Println("sasa", len(visitedNodes))
			startVisualise = false
		}
		rl.BeginDrawing()
		drawGrid(grid)
		animateVisited(visitedNodes)
		rl.ClearBackground(BG_COL)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

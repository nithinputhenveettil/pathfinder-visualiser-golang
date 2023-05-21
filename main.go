package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type node struct {
	row             int32
	col             int32
	isStart         bool
	isFinish        bool
	isVisited       bool
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

func getInitGrid() [][]*node {
	var i, j int32
	var grid [][]*node

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
		}
		grid = append(grid, r)
	}

	grid[1][3].isBarrier = true

	return grid

}

func drawGrid(grid [][]*node) {
	for _, r := range grid {
		for _, n := range r {
			x := n.col * BlockSize
			y := n.row * BlockSize
			rl.DrawLine(x, y, x+BlockSize, y, OT_COL)
			rl.DrawLine(x, y, x, y+BlockSize, OT_COL)
			if n.isStart {
				drawStartNode(x, y)
			}
			if n.isFinish {
				drawEndNode(x, y)
			}
			if n.isBarrier && !n.isStart && !n.isFinish {
				drawBarrierNode(x, y)
			}
		}
	}
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
	rl.DrawRectangle(x, y, BlockSize, BlockSize, rl.Black)
}

func litsenMouseClick(grid [][]*node) {
	if rl.IsMouseButtonPressed(0) {
		points := rl.GetMousePosition()
		fmt.Println(points)
		x := (int32)(points.X / float32(BlockSize))
		y := (int32)(points.Y / float32(BlockSize))
		grid[y][x].isBarrier = !grid[y][x].isBarrier
		fmt.Println(points, x, y)
	}
}

func main() {
	grid := getInitGrid()
	rl.InitWindow(width, length, "Path Finder Visualiser")
	rl.SetTargetFPS(FPS)
	for !rl.WindowShouldClose() {
		litsenMouseClick(grid)
		rl.BeginDrawing()
		drawGrid(grid)
		rl.ClearBackground(BG_COL)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

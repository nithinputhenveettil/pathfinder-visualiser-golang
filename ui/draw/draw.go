package draw

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid/node"
)

var (
	BG_COL = rl.White
	OT_COL = rl.Black

	ST_COL = rl.Red
)

func drawNodeOutline(x, y int32) {
	rl.DrawRectangleLines(x+1, y+1, grid.BlockSize-3, grid.BlockSize-3, OT_COL)
}

func drawStartNode(x, y int32) {
	rl.DrawLineEx(rl.Vector2{X: float32(x + (grid.BlockSize / 4)), Y: float32(y + (grid.BlockSize / 4))}, rl.Vector2{X: float32(x + (grid.BlockSize * 3 / 4)), Y: float32(y + (grid.BlockSize / 2))}, 4.0, ST_COL)
	rl.DrawLineEx(rl.Vector2{X: float32(x + (grid.BlockSize / 4)), Y: float32(y + (grid.BlockSize * 3 / 4))}, rl.Vector2{X: float32(x + (grid.BlockSize * 3 / 4)), Y: float32(y + (grid.BlockSize / 2))}, 4.0, ST_COL)
}

func drawEndNode(x, y int32) {
	rl.DrawCircleV(rl.Vector2{X: float32(x + (grid.BlockSize / 2)), Y: float32(y + (grid.BlockSize / 2))}, 3, ST_COL)
	rl.DrawRing(rl.Vector2{X: float32(x + (grid.BlockSize / 2)), Y: float32(y + (grid.BlockSize / 2))}, 6, 10, 0, 360, 0, ST_COL)
}

func drawBarrierNode(x, y int32) {
	rl.DrawRectangle(x+1, y+1, grid.BlockSize-3, grid.BlockSize-3, OT_COL)
}

func drawVisitedNode(x, y int32) {
	rl.DrawRectangle(x+1, y+1, grid.BlockSize-3, grid.BlockSize-3, rl.SkyBlue)
}

func drawVisitedNodeshortPath(x, y int32, previousVisited *node.Node) {
	// rl.DrawRectangle(x+1, y+1, grid.BlockSize-3, grid.BlockSize-3, rl.Yellow)
	if previousVisited == nil {
		return
	}

	xPrevious := previousVisited.Col * grid.BlockSize
	yPrevious := previousVisited.Row * grid.BlockSize

	rl.DrawLineEx(rl.Vector2{
		X: float32(xPrevious + (grid.BlockSize / 2)),
		Y: float32(yPrevious + (grid.BlockSize / 2)),
	}, rl.Vector2{
		X: float32(x + (grid.BlockSize / 2)),
		Y: float32(y + (grid.BlockSize / 2)),
	}, 4.0, ST_COL)

}

func DrawGrid(v *grid.Visualiser) {
	rl.ClearBackground(BG_COL)
	for _, r := range v.Grid {
		for _, n := range r {
			x := n.Col * grid.BlockSize
			y := n.Row * grid.BlockSize
			drawNodeOutline(x, y)

			if n.IsBarrier && !n.IsStart && !n.IsFinish {
				drawBarrierNode(x, y)
			}
			if n.IsVisited && n.AnimateVisited {
				drawVisitedNode(x, y)
			}
			if n.IsStart {
				drawStartNode(x, y)
			}
			if n.IsFinish {
				drawEndNode(x, y)
			}

		}
	}

	for _, n := range v.ShortPathNodes {
		x := n.Col * grid.BlockSize
		y := n.Row * grid.BlockSize
		if n.IsVisited && n.AnimateShortPath && !n.IsStart {
			drawVisitedNodeshortPath(x, y, n.PreviousVisited)
		}
	}
}

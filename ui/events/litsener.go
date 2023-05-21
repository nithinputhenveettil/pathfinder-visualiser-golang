package events

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid"
)

func LitsenMouseClick(v *grid.Visualiser) {
	if rl.IsMouseButtonPressed(0) {
		points := rl.GetMousePosition()
		if points.X <= 0 || points.X >= float32(grid.Width) || points.Y <= float32(grid.TopPadding) || points.Y-float32(grid.TopPadding) >= float32(grid.Length) {
			return
		}
		x, y := getXY(points)
		if v.Grid[y][x] == v.StartNode {
			v.StartNode.IsStart = false
			v.StartNode = nil
		} else if v.Grid[y][x] == v.EndNode {
			v.EndNode.IsFinish = false
			v.EndNode = nil
		} else {
			v.Grid[y][x].IsBarrier = !v.Grid[y][x].IsBarrier
		}
	}
	if rl.IsMouseButtonReleased(0) {
		points := rl.GetMousePosition()
		if points.X <= 0 || points.X >= float32(grid.Width) || points.Y <= float32(grid.TopPadding) || points.Y-float32(grid.TopPadding) >= float32(grid.Length) {
			return
		}
		x, y := getXY(points)
		if v.StartNode == nil {
			v.StartNode = v.Grid[y][x]
			v.StartNode.IsStart = true
		} else if v.EndNode == nil {
			v.EndNode = v.Grid[y][x]
			v.EndNode.IsFinish = true
		}
		if v.IsDoneWithFirstCycle {
			v.ResetLastVisit()
			v.StartVisualise = true
		}
	}
	if rl.IsMouseButtonDown(1) {
		points := rl.GetMousePosition()
		if points.X <= 0 || points.X >= float32(grid.Width) || points.Y <= float32(grid.TopPadding) || points.Y-float32(grid.TopPadding) >= float32(grid.Length) {
			return
		}
		x, y := getXY(points)
		v.Grid[y][x].IsBarrier = !v.Grid[y][x].IsBarrier
	}
}

func LitsenKeyboardEvents(v *grid.Visualiser) {
	if rl.IsKeyPressed(83) {
		v.StartVisualise = true
	}
	if rl.IsKeyPressed(82) {
		// reset
		v.Reset()
	}
}

func getXY(points rl.Vector2) (int32, int32) {
	x := (int32)(points.X / float32(grid.BlockSize))
	y := (int32)((points.Y - float32(grid.TopPadding)) / float32(grid.BlockSize))
	return x, y
}

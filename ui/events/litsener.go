package events

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid"
)

func LitsenMouseClick(v *grid.Visualiser) {
	if rl.IsMouseButtonPressed(0) {
		x, y := getXY()
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
		x, y := getXY()
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
		x, y := getXY()
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

func getXY() (int32, int32) {
	points := rl.GetMousePosition()
	x := (int32)(points.X / float32(grid.BlockSize))
	y := (int32)((points.Y - float32(grid.TopPadding)) / float32(grid.BlockSize))
	return x, y
}

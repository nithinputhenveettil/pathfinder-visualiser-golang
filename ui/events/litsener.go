package events

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid"
)

func LitsenMouseClick(v *grid.Visualiser) {
	if rl.IsMouseButtonDown(0) {
		points := rl.GetMousePosition()
		x := (int32)(points.X / float32(grid.BlockSize))
		y := (int32)(points.Y / float32(grid.BlockSize))
		v.Grid[y][x].IsBarrier = !v.Grid[y][x].IsBarrier
	}
	if rl.IsMouseButtonPressed(1) {
		points := rl.GetMousePosition()
		x := (int32)(points.X / float32(grid.BlockSize))
		y := (int32)(points.Y / float32(grid.BlockSize))
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

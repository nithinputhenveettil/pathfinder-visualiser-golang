package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/algorithm"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/ui/animate"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/ui/draw"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/ui/events"
	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid"
)

func main() {
	v := &grid.Visualiser{}
	v.Init()
	rl.InitWindow(grid.Width, grid.Length, "Path Finder Visualiser - Dijkstra's Algorithm")
	rl.SetTargetFPS(grid.FPS)
	for !rl.WindowShouldClose() {
		events.LitsenMouseClick(v)
		events.LitsenKeyboardEvents(v)
		if v.StartVisualise {
			v.VisitedNodes = algorithm.Dijkstra(v)
			v.ShortPathNodes = algorithm.GetNodesInShortPath(v)
			v.StartVisualise = false
		}
		rl.BeginDrawing()
		draw.DrawGrid(v)
		if !v.IsDoneWithFirstCycle {
			animate.NextTickAnimate(v)
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

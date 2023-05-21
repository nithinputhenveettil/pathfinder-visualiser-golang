# pathfinder-visualiser-golang

This is golang implementation of pathfinder-visualiser - Dijkstra's Algorithm!
- Inspired by [Clément's Pathfinding Visualizer](https://clementmihailescu.github.io/Pathfinding-Visualizer/)

![pathfinder-visualiser-golang](https://github.com/nithinputhenveettil/pathfinder-visualiser-golang/assets/25578971/36c2a236-5fc3-4d5d-b2a8-16be66813b82)

# Dependencies

This repo uses `raylib` as a dependency. You need to satisfy the requirments of `raylib`. 

Read more from [here](https://github.com/gen2brain/raylib-go#requirements) 

# Usage

- Clone this repo
```sh
go build cmd/pathfinder
```
```sh
./pathfinder
```

# Instructions
- Press 's' to start visualise
- Press 'r' to reset everything
- Use left/right mouse button to add a barrier in grid
- Drag start/end node to re-position it in the grid
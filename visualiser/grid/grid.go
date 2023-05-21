package grid

import (
	"math"

	"github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid/node"
)

var (
	start_row int32 = 10
	start_col int32 = 10

	end_row int32 = 10
	end_col int32 = 40
)

var (
	Width  int32 = 1300
	Length int32 = 650
	FPS    int32 = 120

	BlockSize int32   = 25
	Infinity  float64 = math.Inf(1)
)

type Visualiser struct {
	IsDoneWithFirstCycle bool
	Grid                 [][]*node.Node
	StartVisualise       bool
	StartNode            *node.Node
	EndNode              *node.Node
	VIndex               int
	SIndex               int
	VisitedNodes         []*node.Node
	ShortPathNodes       []*node.Node
}

func (v *Visualiser) Init() {
	var g [][]*node.Node
	var i, j int32

	l := Length / BlockSize
	w := Width / BlockSize

	for i = 0; i < l; i++ {
		r := []*node.Node{}
		for j = 0; j < w; j++ {
			n := &node.Node{
				Row:      i,
				Col:      j,
				IsStart:  j == start_col && i == start_row,
				IsFinish: j == end_col && i == end_row,
				Distance: Infinity,
			}
			r = append(r, n)
			if n.IsStart {
				v.StartNode = n
			}
			if n.IsFinish {
				v.EndNode = n
			}
		}
		g = append(g, r)
	}
	v.Grid = g
}

func (v *Visualiser) Reset() {
	v.Grid = [][]*node.Node{}
	v.StartVisualise = false
	v.StartNode = nil
	v.EndNode = nil
	v.VIndex = 0
	v.SIndex = 0
	v.VisitedNodes = []*node.Node{}
	v.ShortPathNodes = []*node.Node{}
	v.IsDoneWithFirstCycle = false
	v.Init()
}

func (v *Visualiser) ResetLastVisit() {
	v.VIndex = 0
	v.SIndex = 0
	for _, r := range v.Grid {
		for _, n := range r {
			n.PreviousVisited = nil
			n.AnimateShortPath = false
			n.AnimateVisited = false
			n.IsVisited = false
			n.Distance = Infinity
		}
	}
	v.VisitedNodes = []*node.Node{}
	v.ShortPathNodes = []*node.Node{}
}

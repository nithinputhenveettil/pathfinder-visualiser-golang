package animate

import "github.com/nithinputhenveettil/pathfinder-visualiser-golang/visualiser/grid"

func NextTickAnimate(v *grid.Visualiser) {
	if v.VIndex < len(v.VisitedNodes) {
		n := v.VisitedNodes[v.VIndex]
		n.AnimateVisited = true
		v.VIndex += 1
	} else if v.SIndex < len(v.ShortPathNodes) {
		n := v.ShortPathNodes[v.SIndex]
		n.AnimateShortPath = true
		v.SIndex += 1
	} else {
		if !(v.SIndex == 0 || v.VIndex == 0) {
			v.IsDoneWithFirstCycle = true
		}
		return
	}
}

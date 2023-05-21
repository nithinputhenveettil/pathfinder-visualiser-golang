package node

type Node struct {
	Row              int32
	Col              int32
	IsStart          bool
	IsFinish         bool
	IsVisited        bool
	AnimateVisited   bool
	AnimateShortPath bool
	IsBarrier        bool
	PreviousVisited  *Node
	Distance         float64
}

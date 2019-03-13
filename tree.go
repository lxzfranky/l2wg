package l2wg

type nodeType uint8

type node struct {
	path      string
	wildChild bool
	nType     nodeType
	maxParams uint8
	indices   string
	children  []*node
	handle    []Handler
	priority  uint32
}




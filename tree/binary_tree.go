package tree

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

type BinaryTree struct {
	root     *Node
	children []*Node
	height   int
	depth    int
}

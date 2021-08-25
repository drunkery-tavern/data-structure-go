package tree

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

type BinaryTree struct {
	root   *Node
	height int
	depth  int
}

func NewRoot(node *Node) *BinaryTree {
	return &BinaryTree{root: node, height: 0, depth: 0}
}

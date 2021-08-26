package tree

type Node struct {
	binaryTreeEntry
	times  int
	parent *Node
	left   *Node
	right  *Node
}

type BinarySearchTree struct {
	root *Node
}

func (b *BinarySearchTree) Root() interface{} {
	return b.root
}

func NewTreeRoot(node *Node) *BinarySearchTree {
	return &BinarySearchTree{root: node}
}

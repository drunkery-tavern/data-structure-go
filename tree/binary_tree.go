package tree

type BinaryTreeEntry struct {
	Key   uint32
	Value interface{}
}
type IBinaryTree interface {
	Root() interface{}
}

type Node struct {
	data  interface{}
	left  *Node
	right *Node
	//height int
	//depth  int
}

type BinaryTree struct {
	BinaryTreeEntry
	root *Node
}

func (b *BinaryTree) Root() interface{} {
	return b.root.data
}

func NewTreeRoot(node *Node) *BinaryTree {
	return &BinaryTree{root: node}
}

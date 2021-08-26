package tree

type binaryTreeEntry struct {
	Key   uint32
	Value interface{}
}
type IBinaryTree interface {
	Root() interface{}
}

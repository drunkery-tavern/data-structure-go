package binaryTree

type binaryTreeEntry struct {
	Key   uint32
	Value interface{} //节点存储的值
}
type IBinaryTree interface {
	IsNil(interface{}) bool
	Root() interface{}
	Height(interface{}) int
	Search(uint32) interface{}
	Insert(key uint32, value interface{})
	Delete(uint32) interface{}
	Predecessor(interface{}, interface{}) interface{}
	Successor(interface{}, interface{}) interface{}
	LeftRotate(interface{}) interface{}
	RightRotate(interface{}) interface{}
	LeftRightRotate(interface{}) interface{}
	RightLeftRotate(interface{}) interface{}
	Min(interface{}) interface{}
	Max(interface{}) interface{}
	PreOrderTraverse(interface{}, []interface{}) []interface{}
	PostOrderTraverse(interface{}, []interface{}) []interface{}
	InOrderTraverse(interface{}, []interface{}) []interface{}
}

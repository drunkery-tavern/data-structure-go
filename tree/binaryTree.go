package tree

type binaryTreeEntry struct {
	Key   uint32      //中序遍历的节点序号
	Value interface{} //节点存储的值
}
type IBinaryTree interface {
	IsNil(interface{}) bool
	Root() interface{}
	Search(uint32) interface{}
	Insert(key uint32, value interface{})
	Delete(uint32) interface{}
	Predecessor(interface{}, interface{}) interface{}
	Successor(interface{}, interface{}) interface{}
	LeftRotate(interface{}) interface{}
	RightRotate(interface{}) interface{}
	Min() interface{}
	Max() interface{}
	PreOrderTraverse(interface{}, []interface{}) []interface{}
	PostOrderTraverse(interface{}, []interface{}) []interface{}
	InOrderTraverse(interface{}, []interface{}) []interface{}
}

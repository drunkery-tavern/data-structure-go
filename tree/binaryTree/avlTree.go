package binaryTree

import "sync"

type AVLNode struct {
	binaryTreeEntry
	height int
	left   *AVLNode
	right  *AVLNode
}

type AVLTree struct {
	root *AVLNode
	lock sync.Mutex
}

func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

func (t *AVLTree) IsNil(node interface{}) bool {
	panic("implement me")
}

func (t *AVLTree) Root() interface{} {
	panic("implement me")
}

func (t *AVLTree) Search(key uint32) interface{} {
	panic("implement me")
}

func (t *AVLTree) Insert(key uint32, value interface{}) {
	panic("implement me")
}

func (t *AVLTree) Delete(key uint32) interface{} {
	panic("implement me")
}

func (t *AVLTree) Predecessor(node interface{}, root interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) Successor(node interface{}, root interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) LeftRotate(node interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) RightRotate(node interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) LeftRightRotate(node interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) RightLeftRotate(node interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) Min(node interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) Max(node interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) PreOrderTraverse(node interface{}, nodeSlice []interface{}) []interface{} {
	panic("implement me")
}

func (t *AVLTree) PostOrderTraverse(node interface{}, nodeSlice []interface{}) []interface{} {
	panic("implement me")
}

func (t *AVLTree) InOrderTraverse(node interface{}, nodeSlice []interface{}) []interface{} {
	panic("implement me")
}

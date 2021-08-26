package tree

import "sync"

type Node struct {
	binaryTreeEntry
	times  int
	parent *Node
	left   *Node
	right  *Node
}

type BinarySearchTree struct {
	root *Node
	lock sync.Mutex
}

// IsNil 判断节点是否为nil
func (t *BinarySearchTree) IsNil(n interface{}) bool {
	return n == nil
}

// Search 根据key查找节点
func (t *BinarySearchTree) Search(key uint32) interface{} {
	for cur := t.root; cur != nil; {
		if cur.Key == key {
			return cur
		} else if key < cur.Key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}

func (t *BinarySearchTree) Insert(i interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Delete(u uint32) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Predecessor(i interface{}, i2 interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Successor(i interface{}, i2 interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) LeftRotate(i interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) RightRotate(i interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Min(i interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Max(i interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Root() interface{} {
	return t.root
}

func NewTreeRoot() *BinarySearchTree {
	return new(BinarySearchTree)
}

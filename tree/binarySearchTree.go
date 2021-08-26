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

func (t *BinarySearchTree) Insert(n interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Delete(key uint32) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Predecessor(n interface{}, root interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Successor(n interface{}, root interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) LeftRotate(n interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) RightRotate(n interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Min(n interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Max(n interface{}) interface{} {
	panic("implement me")
}

func (t *BinarySearchTree) Root() interface{} {
	return t.root
}

func NewTreeRoot() *BinarySearchTree {
	return new(BinarySearchTree)
}

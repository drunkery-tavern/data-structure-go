package binaryTree

import (
	"math"
	"sync"
)

type AVLNode struct {
	binaryTreeEntry
	height int //节点的高度，用于计算父节点的平衡因子
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

func (t *AVLTree) Height(node interface{}) int {
	if node == nil {
		return 0
	}
	return node.(*AVLNode).height
}

// BalanceFactor 计算平衡因子
func (t *AVLTree) BalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return t.Height(node.left) - t.Height(node.right)
}

func (t *AVLTree) IsNil(node interface{}) bool {
	return node == nil
}

func (t *AVLTree) Root() interface{} {
	return t.root
}

func (t *AVLTree) Search(key uint32) interface{} {
	for cur := t.root; cur != nil; {
		if cur.Key == key {
			return cur
		} else if cur.Key < key {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	return nil
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
	n := node.(*AVLNode)
	nr := n.right
	reConnectNode := nr.left
	nr.left, n.right = n, reConnectNode
	n.height = int(math.Max(float64(t.Height(n.left)), float64(t.Height(n.right))) + 1)
	nr.height = int(math.Max(float64(t.Height(nr.left)), float64(t.Height(nr.right))) + 1)
	return nr
}

func (t *AVLTree) RightRotate(node interface{}) interface{} {
	n := node.(*AVLNode)
	nl := n.left
	reConnectNode := nl.right
	nl.right, n.left = n, reConnectNode
	n.height = int(math.Max(float64(t.Height(n.left)), float64(t.Height(n.right))) + 1)
	nl.height = int(math.Max(float64(t.Height(nl.left)), float64(t.Height(nl.right))) + 1)
	return nl
}

func (t *AVLTree) LeftRightRotate(node interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) RightLeftRotate(node interface{}) interface{} {
	panic("implement me")
}

func (t *AVLTree) Min(node interface{}) interface{} {
	current := node.(*AVLNode)
	for current.left != nil {
		current = current.left
	}
	return current
}

func (t *AVLTree) Max(node interface{}) interface{} {
	current := node.(*AVLNode)
	for current.right != nil {
		current = current.right
	}
	return current
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

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

func (t *BinarySearchTree) Insert(key uint32, value interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	node := new(Node)
	node.Key = key
	node.Value = value

	if t.root == nil {
		t.root = node
	} else {
		target := t.root
		for cur := t.root; cur != nil; {
			target = cur
			if node.Key < cur.Key {
				cur = cur.left
			} else {
				cur = cur.right
			}
		}
		node.parent = target
		if node.Key < target.Key {
			target.left = node
		} else {
			target.right = node
		}
	}
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

func (t *BinarySearchTree) Min() interface{} {
	current := t.root
	for current = t.root; current.left != nil; {
		current = current.left
	}
	return current
}

func (t *BinarySearchTree) Max() interface{} {
	current := t.root
	for current = t.root; current.right != nil; {
		current = current.right
	}
	return current
}

func (t *BinarySearchTree) Root() interface{} {
	return t.root
}

// PreOrderTraverse 前序遍历  根节点 -> 左子树 -> 右子树
func (t *BinarySearchTree) PreOrderTraverse(node interface{}) interface{} {
	n := node.(*Node)
	var nodeSlice []*Node
	if n != nil {
		nodeSlice = append(nodeSlice, n)
		if n.left != nil {
			t.PreOrderTraverse(n.left)
		}
		if n.right != nil {
			t.PreOrderTraverse(n.right)
		}
	}
	return nodeSlice
}

// PostOrderTraverse 后序遍历
func (t *BinarySearchTree) PostOrderTraverse() interface{} {
	panic("implement me")
}

// InOrderTraverse 中序遍历
func (t *BinarySearchTree) InOrderTraverse() interface{} {
	panic("implement me")
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

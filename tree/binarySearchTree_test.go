package tree

import (
	"log"
	"testing"
)

var tree *BinarySearchTree

func init() {
	tree = NewBinarySearchTree()
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	tree.Insert(3, 3)
	tree.Insert(4, 4)
	tree.Insert(5, 5)
}

func TestBinarySearchTree_Delete(t *testing.T) {

}

func TestBinarySearchTree_InOrderTraverse(t *testing.T) {

}

func TestBinarySearchTree_Insert(t *testing.T) {
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	tree.Insert(3, 3)
	tree.Insert(4, 4)
	tree.Insert(5, 5)
}

func TestBinarySearchTree_IsNil(t *testing.T) {

}

func TestBinarySearchTree_LeftRotate(t *testing.T) {

}

func TestBinarySearchTree_Max(t *testing.T) {

}

func TestBinarySearchTree_Min(t *testing.T) {

}

func TestBinarySearchTree_PostOrderTraverse(t *testing.T) {

}

func TestBinarySearchTree_PreOrderTraverse(t *testing.T) {
	traverse := tree.PreOrderTraverse(tree.root)
	ints := traverse.([]*Node)
	for _, v := range ints {
		log.Printf("%v", v)
	}
}

func TestBinarySearchTree_Predecessor(t *testing.T) {

}

func TestBinarySearchTree_RightRotate(t *testing.T) {

}

func TestBinarySearchTree_Root(t *testing.T) {

}

func TestBinarySearchTree_Search(t *testing.T) {

}

func TestBinarySearchTree_Successor(t *testing.T) {

}

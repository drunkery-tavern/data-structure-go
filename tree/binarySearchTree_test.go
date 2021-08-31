package tree

import (
	"log"
	"strconv"
	"testing"
)

var tree *BinarySearchTree

func init() {
	tree = NewBinarySearchTree()
	tree.Insert(8, 8)
	tree.Insert(3, 3)
	tree.Insert(10, 10)
	tree.Insert(1, 1)
	tree.Insert(6, 6)
	tree.Insert(4, 4)
	tree.Insert(7, 7)
	tree.Insert(14, 14)
	tree.Insert(13, 13)
}

func TestBinarySearchTree_Delete(t *testing.T) {

}

func TestBinarySearchTree_InOrderTraverse(t *testing.T) {
	traverse := tree.InOrderTraverse(tree.root, []interface{}{})
	stringify := ""
	for i, v := range traverse {
		if i == len(traverse)-1 {
			stringify += strconv.Itoa(v.(int))
			break
		}
		stringify += strconv.Itoa(v.(int)) + " -> "
	}
	log.Println(stringify)
}

func TestBinarySearchTree_Insert(t *testing.T) {
	tree.Insert(8, 8)
	tree.Insert(3, 3)
	tree.Insert(10, 10)
	tree.Insert(1, 1)
	tree.Insert(6, 6)
	tree.Insert(4, 4)
	tree.Insert(7, 7)
	tree.Insert(14, 14)
	tree.Insert(13, 13)
}

func TestBinarySearchTree_IsNil(t *testing.T) {

}

func TestBinarySearchTree_LeftRotate(t *testing.T) {

}

func TestBinarySearchTree_Max(t *testing.T) {
	log.Println(tree.Max())
}

func TestBinarySearchTree_Min(t *testing.T) {
	log.Println(tree.Min())
}

func TestBinarySearchTree_PostOrderTraverse(t *testing.T) {
	traverse := tree.PostOrderTraverse(tree.root, []interface{}{})
	stringify := ""
	for i, v := range traverse {
		if i == len(traverse)-1 {
			stringify += strconv.Itoa(v.(int))
			break
		}
		stringify += strconv.Itoa(v.(int)) + " -> "
	}
	log.Println(stringify)
}

func TestBinarySearchTree_PreOrderTraverse(t *testing.T) {
	traverse := tree.PreOrderTraverse(tree.root, []interface{}{})
	stringify := ""
	for i, v := range traverse {
		if i == len(traverse)-1 {
			stringify += strconv.Itoa(v.(int))
			break
		}
		stringify += strconv.Itoa(v.(int)) + " -> "
	}
	log.Println(stringify)
}

func TestBinarySearchTree_Predecessor(t *testing.T) {

}

func TestBinarySearchTree_RightRotate(t *testing.T) {

}

func TestBinarySearchTree_Root(t *testing.T) {
	log.Println(tree.Root())
}

func TestBinarySearchTree_Search(t *testing.T) {
	log.Println(tree.Search(5))
	log.Println(tree.Search(8))
}

func TestBinarySearchTree_Successor(t *testing.T) {

}

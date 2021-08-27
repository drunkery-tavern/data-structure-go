package main

import (
	"data-structure-go/tree"
	"log"
)

func main() {

	t := tree.NewBinarySearchTreeRoot()
	log.Println(t.Insert(uint32(1)))
	log.Println(t.Insert(uint32(2)))
	log.Println(t.Insert(uint32(3)))
}

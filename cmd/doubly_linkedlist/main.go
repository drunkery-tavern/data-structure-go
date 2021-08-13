package main

import (
	"data-structure-go/linkedlist"
	"log"
)

func main() {

	linkedList := linkedlist.DoublyLinkedList{}
	list := linkedList.NewDoublyLinkedList()
	list.AddFirst(298)
	list.AddFirst(4344444)
	log.Println(list.ToSlice())

	list.AddLast(333)
	log.Println(list.ToSlice())
}

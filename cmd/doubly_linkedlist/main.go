package main

import "data-structure-go/linkedlist"

func main() {

	linkedList := linkedlist.DoublyLinkedList{}
	list := linkedList.NewDoublyLinkedList()
	list.AddFirst(298)
	println(list.Len())

}

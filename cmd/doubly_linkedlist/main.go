package main

import (
	"data-structure-go/linkedlist/doubly_linkedlist"
	"log"
)

func main() {

	linkedList := doubly_linkedlist.DoublyLinkedList{}
	list := linkedList.NewDoublyLinkedList()
	list.AddFirst(298)
	list.AddFirst(4344444)
	log.Println(list.ToSlice())

	list.AddLast(333)
	log.Println(list.ToSlice())

	err := list.AddByIndex(345, 1)
	if err != nil {
		return
	}

	log.Println(list.ToSlice())

	err = list.AddByIndex(3675, 3)
	if err != nil {
		return
	}

	log.Println(list.ToSlice())
}

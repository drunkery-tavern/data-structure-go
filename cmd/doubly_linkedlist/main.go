package main

import (
	"data-structure-go/linkedlist"
	"log"
)

func main() {

	list := linkedlist.NewDoublyLinkedList()
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

	log.Println(list.GetByIndex(3))

	log.Println(list.RemoveFirst())

	log.Println(list.ToSlice())

	log.Println(list.RemoveLast())

	log.Println(list.ToSlice())

	log.Println(list.RemoveByIndex(1))
	log.Println(list.ToSlice())
}

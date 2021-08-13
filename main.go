package main

import (
	"data-structure-go/linkedlist"
	"log"
)

func main() {
	linkedList := linkedlist.SingleLinkedList{}
	list := linkedList.NewSingleLinkedList()
	list.AddFirst(1)
	list.AddLast(3)
	list.AddLast(7)
	list.AddLast(9)

	err := list.AddByIndex(5, 5)
	if err != nil {
		log.Println(err)
	}

	err = list.AddByIndex(5, 1)
	if err != nil {
		log.Println(err)
	}

	log.Println(list.ToSlice())

	log.Println(list.GetByIndex(1))
	last := list.RemoveLast()
	log.Println(last)
	log.Println(list.ToSlice())

	log.Println(list.RemoveByIndex(1))
	log.Println(list.ToSlice())
}

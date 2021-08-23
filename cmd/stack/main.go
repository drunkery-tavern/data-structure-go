package main

import (
	"data-structure-go/stack"
	"log"
)

func main() {
	s := stack.NewStack()
	s.Push(2)
	s.Push(4)
	s.Push(6)
	log.Println(s.Len())
	//log.Println(s.Pop())
	//log.Println(s.Peek())
	log.Println(s.ToSlice())
}

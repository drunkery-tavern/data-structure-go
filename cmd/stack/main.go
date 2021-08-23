package main

import (
	"data-structure-go/stack"
	"log"
)

func main() {
	st := stack.Stack{}
	s := st.NewStack()
	log.Println(s.ToSlice())
	s.Push(2)
	s.Push(4)
	s.Push(6)
	log.Println(s.Len())
	log.Println(s.Pop())
	log.Println(s.Peek())
}

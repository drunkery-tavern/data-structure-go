package main

import (
	"data-structure-go/queue"
	"log"
)

func main() {
	q := queue.NewQueue()
	q.Push(4)
	q.Push(6)
	q.Push(8)
	q.Push(10)
	log.Println(q.Pop())
	log.Println(q.Len())
	log.Println(q.Peek())
	log.Println(q.ToSlice())
}

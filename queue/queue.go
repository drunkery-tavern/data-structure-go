package queue

import "sync"

type IQueue interface {
	Len() int
	Peek() interface{}
	Push(v interface{})
	Pop() interface{}
	ToSlice() []interface{}
	Any() bool
}

type (
	//Queue 队列
	Queue struct {
		head *Node
		tail *Node
		len  int
		lock sync.Mutex
	}
	// Node 双向链表节点
	Node struct {
		prev *Node
		next *Node
		data interface{}
	}
)

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Peek() interface{} {
	if q.len == 0 {
		return nil
	}
	return q.head.data
}

func (q *Queue) Push(v interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	node := &Node{
		prev: nil,
		next: nil,
		data: v,
	}
	if q.len == 0 {
		q.head = node
		q.tail = q.head
	}
	node.prev = q.tail
	q.tail.next = node
	q.tail = node
	q.len++
}

func (q *Queue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.len == 0 {
		return nil
	}
	head := q.head
	q.head.next.prev = nil
	q.head = q.head.next
	q.len--
	return head.data
}

func (q *Queue) ToSlice() (res []interface{}) {
	if q.len == 0 {
		return
	}
	current := q.head
	for {
		res = append(res, current.data)
		if current.next == nil {
			break
		}
		current = current.next
	}
	return
}

func (q *Queue) Any() bool {
	return q.len > 0
}

func NewQueue() *Queue {
	return &Queue{head: nil, tail: nil, len: 0}
}

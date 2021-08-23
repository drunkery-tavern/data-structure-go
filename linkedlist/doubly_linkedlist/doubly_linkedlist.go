package doubly_linkedlist

import (
	"errors"
	"sync"
)

type IList interface {
	Len() int
	AddFirst(v interface{})
	AddLast(v interface{})
	AddByIndex(v interface{}, index int) error
	GetFirst() *Node
	GetLast() *Node
	GetByIndex(index int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	RemoveByIndex(index int) interface{}
	ToSlice() (res []interface{})
}

type (
	DoublyLinkedList struct {
		header *Node
		tail   *Node
		len    int
		lock   sync.Mutex
	}

	Node struct {
		data interface{}
		next *Node
		prev *Node
	}
)

func (*DoublyLinkedList) NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		header: nil,
		tail:   nil,
		len:    0,
	}
}

func (l *DoublyLinkedList) Len() int {
	return l.len
}

func (l *DoublyLinkedList) AddFirst(v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	defer func() { l.len++ }()
	node := &Node{
		data: v,
		next: nil,
		prev: nil,
	}
	if l.len == 0 {
		l.tail = node
	} else {
		l.header.prev = node
	}
	node.next = l.header
	l.header = node
}

func (l *DoublyLinkedList) AddLast(v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	defer func() { l.len++ }()
	node := &Node{
		data: v,
		next: nil,
		prev: nil,
	}

	if l.len == 0 {
		l.header = node
	} else {
		l.tail.next = node
	}
	node.prev = l.tail
	l.tail = node
}

func (l *DoublyLinkedList) AddByIndex(v interface{}, index int) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	//索引越界
	if index < 0 || index >= l.len {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		l.AddFirst(v)
		return nil
	}
	if index == l.len {
		l.AddLast(v)
		return nil
	}
	//查找插入位置的相关节点
	mid := l.len / 2
	current := &Node{}
	if index <= mid {
		//从头结点开始遍历
		current = l.header
		for i := 1; i < mid; i++ {
			if i == index {
				break
			}
			current = current.next
		}
	}
	if index > mid {
		//从尾节点开始遍历
		current = l.tail
		for i := l.len; i > mid; i-- {
			if i == index {
				break
			}
			current = current.prev
		}
	}
	newNode := &Node{
		data: v,
		next: nil,
		prev: nil,
	}
	next := current.next
	current.next = newNode
	newNode.prev = current
	newNode.next = next
	next.prev = newNode
	l.len++
	return nil
}

func (l *DoublyLinkedList) GetFirst() *Node {
	return l.header
}

func (l *DoublyLinkedList) GetLast() *Node {
	return l.tail
}

func (l *DoublyLinkedList) GetByIndex(index int) interface{} {
	//索引越界
	if index < 0 || index >= l.len {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		return l.GetFirst().data
	}
	if index == l.len {
		return l.GetLast().data
	}
	current := l.header
	for i := 0; i < l.len; i++ {
		if i == index {
			return current.data
		}
		current = current.next
	}
	return nil
}

func (l *DoublyLinkedList) RemoveFirst() interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.len == 0 {
		return nil
	}
	header := l.header
	header.next.prev = nil
	l.header = header.next
	header.next = nil
	l.len--
	return header.data
}

func (l *DoublyLinkedList) RemoveLast() interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.len == 0 {
		return nil
	}
	tail := l.tail
	tail.prev.next = nil
	l.tail = tail.prev
	tail.prev = nil
	l.len--
	return tail.data
}

func (l *DoublyLinkedList) RemoveByIndex(index int) interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.len == 0 {
		return nil
	}
	//索引越界
	if index < 0 || index >= l.len {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		return l.RemoveFirst()
	}
	if index == l.len-1 {
		return l.RemoveLast()
	}
	//可使用二分法减少循环次数
	current := l.header
	for i := 0; i < l.len; i++ {
		if i == index {
			current.prev.next = current.next
			current.next.prev = current.prev
			current.prev = nil
			current.next = nil
			l.len--
			return current.data
		}
		current = current.next
	}

	return nil
}

func (l *DoublyLinkedList) ToSlice() (res []interface{}) {
	if l.len == 0 {
		return
	}
	current := l.header
	for i := 0; i < l.len; i++ {
		res = append(res, current.data)
		current = current.next
	}
	return
}

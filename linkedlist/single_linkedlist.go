package linkedlist

import (
	"errors"
)

type SingleLinkedList struct {
	header *Node
	tail   *Node
	len    int
}

type Node struct {
	Data interface{}
	next *Node
}

// NewSingleLinkedList 初始化
func (*SingleLinkedList) NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		header: nil,
		tail:   nil,
		len:    0,
	}
}

// Len 获取长度
func (l *SingleLinkedList) Len() int {
	return l.len
}

// AddFirst 头部插入
func (l *SingleLinkedList) AddFirst(v interface{}) {
	defer func() { l.len++ }()
	node := &Node{
		Data: v,
		next: nil,
	}

	//当链表没有元素时，也为尾节点
	if l.len == 0 {
		l.tail = node
	}
	node.next = l.header
	l.header = node
}

// AddLast 尾部插入
func (l *SingleLinkedList) AddLast(v interface{}) {
	defer func() { l.len++ }()
	node := &Node{
		Data: v,
		next: nil,
	}
	if l.len == 0 {
		l.header = node
	}
	l.tail.next = node
	l.tail = node
}

// AddByIndex 指定位置插入
func (l *SingleLinkedList) AddByIndex(v interface{}, index uint) error {
	//索引越界
	if index >= uint(l.len) {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		l.AddFirst(v)
		return nil
	}
	if index == uint(l.len) {
		l.AddLast(v)
		return nil
	}

	index -= 1
	prev := l.header
	for i := 0; i < l.len; i++ {
		if uint(i) == index {
			node := &Node{Data: v, next: nil}
			node.next = prev.next
			prev.next = node
			l.len++
			return nil
		}
		prev = prev.next
	}
	return nil
}

// GetFirst 获取头节点
func (l *SingleLinkedList) GetFirst() *Node {
	return l.header
}

// GetLast 获取尾节点
func (l *SingleLinkedList) GetLast() *Node {
	return l.tail
}

func (l *SingleLinkedList) ToSlice() (res []interface{}) {
	current := l.header
	for i := 0; i < l.len; i++ {
		res = append(res, current.Data)
		current = current.next
	}
	return
}

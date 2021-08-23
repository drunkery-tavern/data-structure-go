package single_linkedlist

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
	SingleLinkedList struct {
		header *Node
		tail   *Node
		len    int
		lock   sync.Mutex
	}

	Node struct {
		data interface{}
		next *Node
	}
)

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
	l.lock.Lock()
	defer l.lock.Unlock()
	defer func() { l.len++ }()
	node := &Node{
		data: v,
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
	l.lock.Lock()
	defer l.lock.Unlock()
	defer func() { l.len++ }()
	node := &Node{
		data: v,
		next: nil,
	}
	if l.len == 0 {
		l.header = node
	}
	l.tail.next = node
	l.tail = node
}

// AddByIndex 指定位置插入
func (l *SingleLinkedList) AddByIndex(v interface{}, index int) error {
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

	index -= 1
	prev := l.header
	for i := 0; i < l.len; i++ {
		if i == index {
			node := &Node{data: v, next: nil}
			node.next = prev.next
			prev.next = node
			l.len++
			return nil
		}
		prev = prev.next
	}
	return errors.New("add node error")
}

// GetFirst 获取头节点
func (l *SingleLinkedList) GetFirst() *Node {
	return l.header
}

// GetLast 获取尾节点
func (l *SingleLinkedList) GetLast() *Node {
	return l.tail
}

// GetByIndex 根据索引获取值
func (l *SingleLinkedList) GetByIndex(index int) interface{} {
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

// RemoveFirst 移除头结点
func (l *SingleLinkedList) RemoveFirst() interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.len == 0 {
		return nil
	}
	data := l.header.data
	l.header = l.header.next
	l.len--
	return data
}

// RemoveLast 移除尾节点
func (l *SingleLinkedList) RemoveLast() interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.len == 0 {
		return nil
	}
	data := l.tail.data
	current := l.header
	for i := 0; i < l.len-1; i++ {
		current = current.next
	}
	current.next = nil
	l.tail = current
	l.len--
	return data
}

// RemoveByIndex 根据索引移除节点
func (l *SingleLinkedList) RemoveByIndex(index int) interface{} {
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

	prev := l.header
	for i := 1; i < index; i++ {
		prev = prev.next
	}
	current := prev.next
	prev.next = prev.next.next
	current.next = nil
	l.len--
	return current.data
}

// ToSlice 转换成切片
func (l *SingleLinkedList) ToSlice() (res []interface{}) {
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

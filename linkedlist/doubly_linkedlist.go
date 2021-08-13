package linkedlist

type DoublyLinkedList struct {
	header *DoublyNode
	tail   *DoublyNode
	len    int
}

type DoublyNode struct {
	Data interface{}
	next *DoublyNode
	prev *DoublyNode
}

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
	defer func() { l.len++ }()
	node := &DoublyNode{
		Data: v,
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
	panic("implement me")
}

func (l *DoublyLinkedList) AddByIndex(v interface{}, index int) error {
	panic("implement me")
}

func (l *DoublyLinkedList) GetFirst() *Node {
	panic("implement me")
}

func (l *DoublyLinkedList) GetLast() *Node {
	panic("implement me")
}

func (l *DoublyLinkedList) GetByIndex(index int) interface{} {
	panic("implement me")
}

func (l *DoublyLinkedList) RemoveFirst() interface{} {
	panic("implement me")
}

func (l *DoublyLinkedList) RemoveLast() interface{} {
	panic("implement me")
}

func (l *DoublyLinkedList) RemoveByIndex(index int) interface{} {
	panic("implement me")
}

func (l *DoublyLinkedList) ToSlice() (res []interface{}) {
	panic("implement me")
}

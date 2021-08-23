package stack

import "sync"

type IStack interface {
	Len() int
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	ToSlice() []interface{}
}

type (
	Stack struct {
		top  *Node
		len  int
		lock sync.Mutex
	}

	Node struct {
		data interface{}
		next *Node
	}
)

func (s *Stack) NewStack() *Stack {
	return &Stack{
		top: nil,
		len: 0,
	}
}

func (s *Stack) Len() int {
	return s.len
}

// Push 进栈
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	node := &Node{
		data: v,
		next: s.top,
	}
	s.top = node
	s.len++
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.len == 0 {
		return nil
	}
	value := s.top.data
	s.top = s.top.next
	s.len--
	return value
}

// Peek 查看栈顶元素
func (s *Stack) Peek() interface{} {
	if s.len == 0 {
		return nil
	}
	return s.top.data
}

// ToSlice 转换成切片
func (s *Stack) ToSlice() (res []interface{}) {
	if s.len == 0 {
		return
	}
	current := s.top
	for {
		if current.next == nil {
			break
		}
		res = append(res, current.data)
		current = current.next
	}
	return
}

package linkedlist

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

package linkedList

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value any
	next  *Node
	prev  *Node
}
type LinkedList struct {
	head Node // 头节点
	len  int  // list长度
}

// Init initializes or clears list l.
func (l *LinkedList) Init() *LinkedList {
	l.head.next = &l.head
	l.head.prev = &l.head
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *LinkedList {
	l := LinkedList{}
	return l.Init()
}

// func (l *LinkedList) lazyInit() {
// 	if l.head.next == nil {
// 		l.Init()
// 	}
// }

func (l *LinkedList) insert(e, at *Node) *Node {
	e.next = at.next
	e.prev = at
	at.next.prev = e
	at.next = e

	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *LinkedList) insertValue(v any, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

// Front returns the first element of list l or nil if the list is empty.
func (l *LinkedList) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *LinkedList) PushBack(v any) *Node {
	// l.lazyInit()
	return l.insertValue(v, l.head.prev)
}

func (l *LinkedList) PushFront(v any) *Node {
	// l.lazyInit()
	return l.insertValue(v, &l.head)
}

func (l *LinkedList) Len() int { return l.len }

func (l *LinkedList) Display() string {
	node := l.Front()
	length := l.len
	var val_list []string
	for length != 0 {

		switch node.Value.(type) {
		case int:
			val_list = append(val_list, strconv.Itoa(node.Value.(int)))
		case string:
			val_list = append(val_list, node.Value.(string))
		default:
			val_list = append(val_list, node.Value.(string))
		}

		node = node.next
		length--
	}

	result := strings.Join(val_list, "->")

	fmt.Println(result)

	return result
}

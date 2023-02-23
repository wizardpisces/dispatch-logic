package linkedList

import (
	"fmt"
	"strconv"
	"strings"
)

type Node[V any] struct {
	Value V
	next  *Node[V]
	prev  *Node[V]
	list  *LinkedList[V]
}
type LinkedList[V any] struct {
	head Node[V] // 头节点
	len  int     // list长度
}

// Init initializes or clears list l.
func (l *LinkedList[V]) Init() *LinkedList[V] {
	l.head.next = &l.head
	l.head.prev = &l.head
	l.len = 0
	return l
}

// New returns an initialized list.
func New[V any]() *LinkedList[V] {
	l := LinkedList[V]{}
	return l.Init()
}

// func (l *LinkedList) lazyInit() {
// 	if l.head.next == nil {
// 		l.Init()
// 	}
// }

func (l *LinkedList[V]) insert(e, at *Node[V]) *Node[V] {
	e.next = at.next
	e.prev = at
	at.next.prev = e
	at.next = e

	e.list = l

	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *LinkedList[V]) insertValue(v V, at *Node[V]) *Node[V] {
	return l.insert(&Node[V]{Value: v}, at)
}

func (l *LinkedList[V]) InsertAfter(v V, at *Node[V]) *Node[V] {
	return l.insert(&Node[V]{Value: v}, at)
}

func (l *LinkedList[V]) InsertBefore(v V, at *Node[V]) *Node[V] {
	return l.insert(&Node[V]{Value: v}, at.prev)
}

// Front returns the first element of list l or nil if the list is empty.
func (l *LinkedList[V]) Front() *Node[V] {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

func (l *LinkedList[V]) Back() *Node[V] {
	if l.len == 0 {
		return nil
	}
	return l.head.prev
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *LinkedList[V]) PushBack(v V) *Node[V] {
	// l.lazyInit()
	return l.insertValue(v, l.head.prev)
}

func (l *LinkedList[V]) PushFront(v V) *Node[V] {
	// l.lazyInit()
	return l.insertValue(v, &l.head)
}

// Next returns the next list element or nil.
func (node *Node[V]) Next() *Node[V] {
	if p := node.next; p != &node.list.head {
		return p
	}
	return nil
}

// remove removes e from its list, decrements l.len
func (l *LinkedList[V]) remove(e *Node[V]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}

func (l *LinkedList[V]) Remove(e *Node[V]) V {
	l.remove(e)
	return e.Value
}

func (l *LinkedList[V]) Len() int { return l.len }

func (l *LinkedList[V]) Display() string {
	var val_list []string

	for node := l.Front(); node != nil; node = node.Next() {
		switch any(node.Value).(type) {
		case int:
			val_list = append(val_list, strconv.Itoa(any(node.Value).(int)))
		case string:
			val_list = append(val_list, any(node.Value).(string))
		default:
			val_list = append(val_list, any(node.Value).(string))
		}
	}

	result := strings.Join(val_list, "->")

	fmt.Println(result)

	return result
}

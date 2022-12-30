package linkedlist

type Node struct {
	Value      interface{}
	next, prev *Node
}
type LinkedList struct {
	head Node // 头节点
	len  int  // list长度
	// tail Node
}

// Init initializes or clears list l.
func (l *LinkedList) Init() *LinkedList {
	l.head.next = &l.head
	l.head.prev = &l.head
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *LinkedList { return new(LinkedList).Init() }

func (l *LinkedList) lazyInit() {
	if l.head.next == nil {
		l.Init()
	}
}

func (l *LinkedList) insert(e, at *Node) *Node {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	// e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *LinkedList) insertValue(v any, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *LinkedList) PushBack(v any) *Node {
	l.lazyInit()
	return l.insertValue(v, l.head.prev)
}

func (l *LinkedList) Len() int { return l.len }

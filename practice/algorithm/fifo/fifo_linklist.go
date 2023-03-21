package fifo

import (
	"sync"
)

// Queue is a FIFO thread-safe queue
type Queue struct {
	head  *node
	tail  *node
	size  int
	mutex sync.Mutex
}

// node is a single-linked list node for the queue
type node struct {
	value interface{}
	next  *node
}

// NewQueue creates a new empty queue and returns it
func NewQueue() *Queue {
	return &Queue{}
}

// Add adds an item at the end of the queue
func (q *Queue) Add(item interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	n := &node{value: item}
	if q.tail == nil { // empty queue
		q.head = n
		q.tail = n
	} else { // non-empty queue
		q.tail.next = n // link new node to tail
		q.tail = n      // update tail pointer to new node
	}
	q.size++
}

// Next pops an item from the front of the queue and returns it.
// Returns nil if the queue is empty.
func (q *Queue) Next() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.head == nil { // empty queue
		return nil
	}
	n := q.head        // get head node
	q.head = n.next    // update head pointer to next node
	if q.head == nil { // last node popped
		q.tail = nil // reset tail pointer to nil as well
	}
	n.next = nil   // unlink popped node from next node (optional)
	q.size--       // decrement size by one (optional)
	return n.value // return popped value

}

// Size returns the number of items in the queue (optional)
func (q *Queue) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.size

}

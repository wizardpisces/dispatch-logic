package fifo

import (
	"fmt"
	"sync"
)

type Queue3 struct {
	items []interface{}
	lock  sync.Mutex
}

func NewQueue3() *Queue3 {
	return &Queue3{}
}

func (q *Queue3) Enqueue(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue3) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func Test_fifo_slice() {
	q := NewQueue3()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Enqueue:", i)
			q.Enqueue(i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			item := q.Dequeue()
			fmt.Println("Dequeue:", item)
		}
	}()

	select {}
}

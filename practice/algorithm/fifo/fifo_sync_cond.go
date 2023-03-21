package fifo

// 生产者，消费者模式
import (
	"fmt"
	"sync"
)

type Queue2 struct {
	items []interface{}
	lock  sync.Mutex
	cond  *sync.Cond
}

func NewQueue2() *Queue2 {
	q := &Queue2{}
	q.cond = sync.NewCond(&q.lock)
	return q
}

func (q *Queue2) Enqueue(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
	q.cond.Signal()
}

func (q *Queue2) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait()
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func Test_fifo_sync_cond() {
	q := NewQueue2()

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

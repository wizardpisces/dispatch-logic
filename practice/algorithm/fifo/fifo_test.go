package fifo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fifo_linklist(t *testing.T) {
	// Create a new queue and add some items to it.
	queue := NewQueue()
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	assert.Equal(t, 3, queue.Size())
	// fmt.Println("Size:", queue.Size()) // Size: 3

	for i := 0; i < 4; i++ {
		item := queue.Next()
		fmt.Println("Item:", item) // Item: 1, Item: 2, Item: 3, Item: <nil>
	}
}

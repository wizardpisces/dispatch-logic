package linkedList

import (
	"container/list"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	assert.Equal(t, l.Len(), 4)
	assert.Equal(t, l.Front().Value, 1)
	assert.Equal(t, l.Back().Value, 4)

	var builder strings.Builder

	for e := l.Front(); e != nil; e = e.Next() {
		// str := fmt.Sprintf("%v", e.Value) // slower
		str := ""
		switch e.Value.(type) {
		case int:
			str = strconv.Itoa(e.Value.(int))
		}
		builder.WriteString(str)
	}

	assert.Equal(t, "1234", builder.String())
	l.Remove(e1)

	assert.Equal(t, l.Len(), 3)
}

func TestLinkedList(t *testing.T) {
	l := New[int]()
	// l.PushBack("2")
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, l.Len(), 2)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	e4 := l.PushBack(4)
	l.PushFront(1)
	e6 := l.InsertAfter(6, e4)
	l.InsertBefore(5, e6)
	result := l.Display()
	assert.Equal(t, "1->2->3->4->5->6", result)
	l.Remove(e4)
	assert.Equal(t, "1->2->3->5->6", l.Display())
}

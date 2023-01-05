package linkedList

import (
	"container/list"
	"fmt"
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

	length := l.Len()
	var builder strings.Builder
	node := l.Front()
	for length != 0 {
		str := fmt.Sprintf("%v", node.Value)
		builder.WriteString(str)
		node = node.Next()
		length--
	}
	assert.Equal(t, "1234", builder.String())
}

func TestLinkedList(t *testing.T) {
	l := New()
	l.PushBack("1")
	l.PushBack("2")
	assert.Equal(t, l.Len(), 2)
	assert.Equal(t, "1", l.head.Value)
	assert.Equal(t, "2", l.head.next.Value)

}
func TestDisplay(t *testing.T) {
	l := New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(4)
	result := l.Display()
	assert.Equal(t, "4->1->2->3", result)
}

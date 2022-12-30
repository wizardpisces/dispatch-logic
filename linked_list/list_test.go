package linkedlist

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_list(t *testing.T) {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	assert.Equal(t, l.Len(), 4)
}

func Test_linkedlist(t *testing.T) {
	l := New()
	l.PushBack(4)
	assert.Equal(t, l.Len(), 1)
}

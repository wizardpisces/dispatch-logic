package prune_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type C int

const (
	C0 C = iota
	C1
	C2
)

const (
	D0 = iota
	D1
)

func Test_PruneTree(t *testing.T) {
	assert.Equal(t, []Digit{nil, 1, 0, 0, 1}, InOrderTraverseTree(ConstructTree([]Digit{1, nil, 0, 0, 1})))
	assert.Equal(t, []Digit{0, 0, 0, 1, 0, 1, 1}, InOrderTraverseTree(ConstructTree([]Digit{1, 0, 1, 0, 0, 0, 1})))

	assert.Equal(t, []Digit{1, nil, 0, 0, 1}, SerializeTree(ConstructTree([]Digit{1, nil, 0, 0, 1})))

	assert.Equal(t, []Digit{1, nil, 0, nil, 1}, SerializeTree(PruneTree(ConstructTree([]Digit{1, nil, 0, 0, 1}))))
	assert.Equal(t, []Digit{1, nil, 1, nil, 1}, SerializeTree(PruneTree(ConstructTree([]Digit{1, 0, 1, 0, 0, 0, 1}))))

	// assert.Equal(t, 0, Zero)
	// assert.Equal(t, 0, C0)
	// assert.Equal(t, 0, D0)
}

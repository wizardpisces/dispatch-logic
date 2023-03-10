package prune_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wizardpisces/dispatch-logic/structures"
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

	assert.Equal(t, []any{1, nil, 0, nil, 1}, structures.TransformToVal(structures.SerializeTree(PruneTree(structures.ConstructTree(structures.TransformToPointer[int]([]any{1, nil, 0, 0, 1}))))))
	assert.Equal(t, []any{1, nil, 1, nil, 1}, structures.TransformToVal(structures.SerializeTree(PruneTree(structures.ConstructTree(structures.TransformToPointer[int]([]any{1, 0, 1, 0, 0, 0, 1}))))))

	assert.Equal(t, 0, C0)
	assert.Equal(t, 0, D0)
}

package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tree(t *testing.T) {
	assert.Equal(t, []any{nil, 1, 0, 0, 1}, TransformToVal(InOrderTraverseTree(ConstructTree(TransformToPointer[int]([]any{1, nil, 0, 0, 1})))))
	assert.Equal(t, []any{0, 0, 0, 1, 0, 1, 1}, TransformToVal(InOrderTraverseTree(ConstructTree(TransformToPointer[int]([]any{1, 0, 1, 0, 0, 0, 1})))))
	assert.Equal(t, []any{1, nil, 0, 0, 1}, TransformToVal(SerializeTree(ConstructTree(TransformToPointer[int]([]any{1, nil, 0, 0, 1})))))

}

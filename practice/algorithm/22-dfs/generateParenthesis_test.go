package generateParenthesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateParenthesis(t *testing.T) {
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, fn(3))
	assert.Equal(t, []string{"()"}, fn(1))
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, fn2(3))
	assert.Equal(t, []string{"()"}, fn2(1))
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, nonRecursive(3))
	assert.Equal(t, []string{"()"}, nonRecursive(1))
	// assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
	assert.Equal(t, []string{"()()()", "()(())", "(())()", "(()())", "((()))"}, generateParenthesis(3))
	assert.Equal(t, []string{"()"}, generateParenthesis(1))

}

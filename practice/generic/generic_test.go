package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneric(t *testing.T) {

	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// When invoking generic functions, we can often rely
	// on _type inference_. Note that we don't have to
	// specify the types for `K` and `V` when
	// calling `MapKeys` - the compiler infers them
	// automatically.
	assert.Equal(t, []int{1, 2, 4}, MapKeys(m))

	// ... though we could also specify them explicitly.
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	assert.Equal(t, []int{10, 13, 23}, lst.GetAll())
}

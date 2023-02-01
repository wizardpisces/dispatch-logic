package tree

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testCase struct {
	name string
	in   []int
	out  []int
}

type mergeSortSuite struct {
	suite.Suite
}

func constructBST(keys []int) *BinarySearchTree {
	bst := BinarySearchTree{}
	for _, key := range keys {
		bst.InsertNode(key, strconv.Itoa(key))
	}
	return &bst
}

func (s mergeSortSuite) Test_Operation() {
	bst := constructBST([]int{2, 3, 1, 6, 5})
	s.Equal("1->2->3->5->6", bst.Display())
	s.Equal([]int{1, 2, 3, 5, 6}, bst.InOrderTraverseTree())
	s.Equal(1, bst.MinNode().key)
	s.Equal(6, bst.MaxNode().key)
	s.Equal(true, bst.SearchNode(1))
	s.Equal(true, bst.SearchNode(2))
	s.Equal(false, bst.SearchNode(7))

	node := bst.DeleteNode(3)
	s.Equal("1->2->5->6", bst.Display())
	s.Equal([]int{1, 2, 5, 6}, bst.InOrderTraverseTree())
	s.Equal(2, node.key)
}

func (s mergeSortSuite) Test_InOrder() {
	cases := []testCase{
		{
			name: "bst in order traverse",
			in:   []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, c := range cases {
		s.Run(c.name, func() {
			s.Equal(c.out, constructBST(c.in).InOrderTraverseTree())
		})
	}
}

func Test(t *testing.T) {
	suite.Run(t, new(mergeSortSuite))
}

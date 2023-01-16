package sort

import (
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

func (s mergeSortSuite) Test_MergeSort() {
	cases := []testCase{
		{
			name: "ascending sorting func",
			in:   []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, c := range cases {
		s.Run(c.name, func() {
			s.Equal(c.out, MergeSort(c.in))
		})
	}
	// unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	// sorted := MergeSort(unsorted)

	// assert.Equal(t, sorted, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestMergeSort(t *testing.T) {
	suite.Run(t, new(mergeSortSuite))
}

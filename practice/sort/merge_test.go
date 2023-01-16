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

// type test2 struct {
// 	name string
// }
// type test3 struct {
// 	name string
// 	age  int
// }
// type test1 struct {
// 	test2
// 	test3
// }

// func (t test2) getName() string {
// 	return t.name
// }
// func (t test3) getName() string {
// 	return t.name
// }

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
}

func TestMergeSort(t *testing.T) {
	// test := new(test1)
	// test.name = "name"
	// test.age = 1

	suite.Run(t, new(mergeSortSuite))
}

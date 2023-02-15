package array_intersection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mergeSortSuite struct {
	suite.Suite
}

// 两个数组的交集
// 给定两个数组，编写一个函数来计算它们的交集。

// 示例 1:

// 输入: nums1 = [1,2,2,1], nums2 = [2,2]

// 输出: [2,2]

// 示例 2:

// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]

// 输出: [4,9]

type input struct {
	arr1, arr2 []int
}

type testCase struct {
	input
	out []int
}

func (s mergeSortSuite) Test() {
	cases := []testCase{
		{
			input: input{
				arr1: []int{1, 2, 2, 1},
				arr2: []int{2, 2},
			},
			out: []int{2, 2},
		},
		{
			input: input{
				arr1: []int{4, 9, 5},
				arr2: []int{9, 4, 9, 8, 4},
			},
			out: []int{9, 4},
		},
	}
	for _, c := range cases {
		// sort.Ints(c.out)
		s.Equal(c.out, ArrayIntersection(c.input.arr1, c.input.arr2))
	}

}

func (s mergeSortSuite) Test2() {
	sortedCases := []testCase{
		{
			input: input{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{2, 3},
			},
			out: []int{2, 3},
		},
		{
			input: input{
				arr1: []int{4, 9, 15},
				arr2: []int{3, 4, 9, 15, 40},
			},
			out: []int{4, 9, 15},
		},
	}
	for _, c := range sortedCases {
		s.Equal(c.out, ArrayIntersectionSorted(c.input.arr1, c.input.arr2))
	}

}

func Test(t *testing.T) {
	suite.Run(t, new(mergeSortSuite))
}

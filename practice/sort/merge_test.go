package sort

import (
	"math/rand"
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
}
func (s mergeSortSuite) Test_MergeSortParallel() {
	cases := []testCase{
		{
			name: "ascending sorting func",
			in:   []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, c := range cases {
		s.Run(c.name, func() {
			s.Equal(c.out, MergeSortParallel(c.in))
		})
	}
}

func TestMergeSort(t *testing.T) {
	suite.Run(t, new(mergeSortSuite))
}

func BenchmarkMergeSort(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	MergeSort(generateIntegers(Threshold + 5))
	// }
}
func BenchmarkMergeSortParallel(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	MergeSortParallel(generateIntegers(Threshold + 5))
	// }
}

func generateIntegers(n int) []int {
	var A []int
	for i := 0; i < n; i++ {
		A = append(A, rand.Intn(n))
	}
	return A
}

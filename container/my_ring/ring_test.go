package my_ring

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	// Create two rings, r and s, of size 2
	r := ring.New(2)
	s := ring.New(2)

	// Get the length of the ring
	lr := r.Len()
	ls := s.Len()

	// Initialize r with 0s
	for i := 0; i < lr; i++ {
		r.Value = 0
		r = r.Next()
	}

	// Initialize s with 1s
	for j := 0; j < ls; j++ {
		s.Value = 1
		s = s.Next()
	}

	// Link ring r and ring s
	rs := r.Link(s)

	// Iterate through the combined ring and print its contents
	rs.Do(func(p any) {
		fmt.Println(p.(int))
	})
}

func TestMyRing(t *testing.T) {
	s := 1
	s ^= 1<<'\t' | 1<<'\n'
	fmt.Println("Hello, 世界", s)
}

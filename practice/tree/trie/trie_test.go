package trie

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mergeSortSuite struct {
	suite.Suite
}

func (s mergeSortSuite) Test() {
	trie := New()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	// wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
	// 	"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for _, word := range words {
		s.Equal(true, trie.Find(word))
	}
	s.Equal(false, trie.Find("sam1"))
}

func Test(t *testing.T) {
	suite.Run(t, new(mergeSortSuite))
}

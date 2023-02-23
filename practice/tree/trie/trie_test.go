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

// func (s mergeSortSuite) TestArrayTrie() {
// 	arrayTrie := NewArray()
// 	list1 := []string{"path", "to", "test1"}
// 	list2 := []string{"path", "to", "test2"}
// 	list3 := []string{"path2", "to", "test1"}
// 	arrayTrie.Insert(list1)
// 	arrayTrie.Insert(list2)
// 	arrayTrie.Insert(list3)

// 	for _, list := range [][]string{list1, list2, list3} {
// 		s.Equal(true, arrayTrie.Find(list))
// 	}

// 	s.Equal(false,arrayTrie.Find([]string{"path"}))
// }

func (s mergeSortSuite) TestParsePattern() {
	s.Equal(parsePattern("/p/*name/*"), []string{"p", "*name"})
	s.Equal(parsePattern("/p/*"), []string{"p", "*"})
	s.Equal(parsePattern("/p/:name"), []string{"p", ":name"})
}

func Test(t *testing.T) {
	suite.Run(t, new(mergeSortSuite))
}

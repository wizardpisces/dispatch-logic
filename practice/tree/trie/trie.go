package trie

type TrieNode struct {
	children map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func newNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func New() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) Insert(str string) *Trie {
	trieNode := t.root

	insert(trieNode, str)

	return t
}

func insert(node *TrieNode, str string) {
	currentNode := node
	for _, ch := range str { // 可以解决乱码问题
		// for i := 0; i < len(str); i++ {
		if currentNode.children[ch] == nil {
			// if currentNode.match(str) {
			currentNode.children[ch] = newNode()
		}
		currentNode = currentNode.children[ch]
	}
}

// func (t *TrieNode) match(part rune) bool {
// 	return false
// }

func (t *Trie) Find(str string) bool {
	return find(t.root, str)
}

func find(node *TrieNode, str string) bool {
	currentNode := node
	for _, ch := range str {
		if currentNode.children[ch] == nil {
			return false
		}
		currentNode = currentNode.children[ch]
	}
	return true
}

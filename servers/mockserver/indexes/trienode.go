package indexes

//TrieNode represents a node in a trie
type TrieNode struct {
	entities *EntityIdentifiers
	children map[rune]*TrieNode
}

//NewTrieNode creates a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{NewEntityIdentifiers(), make(map[rune]*TrieNode)}
}

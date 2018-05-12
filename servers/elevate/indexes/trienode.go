package indexes

//TrieNode represents a node in a trie
type TrieNode struct {
	users    *UserIdentifiers
	children map[rune]*TrieNode
}

//NewTrieNode creates a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{NewUserIdentifiers(), make(map[rune]*TrieNode)}
}

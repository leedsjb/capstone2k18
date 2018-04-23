package indexes

import "testing"

//TestNewNode tests TrieNode's NewNode function
func TestNewNode(t *testing.T) {
	//Test if a struct, which we'll assume to be a TrieNode, is created
	node := NewTrieNode()
	if node == nil {
		t.Error("NewTrieNode() does not return a TrieNode")
	}
	//Test whether the children map is initialized
	if node.children == nil {
		t.Error("TrieNode's children map must be initialized")
	}
}

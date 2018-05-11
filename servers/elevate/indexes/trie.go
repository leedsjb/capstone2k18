package indexes

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
)

// Trie ...
type Trie struct {
	root *TrieNode
	mx   sync.RWMutex
}

// NewTrie ...
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// AddEntity ...
func (trie *Trie) AddEntity(key string, entityID int) error {
	if len(key) == 0 {
		return fmt.Errorf("Please provide a key")
	}

	// Obtain an exclusive lock
	trie.mx.Lock()
	// Let current node = root node
	current := trie.root

	for _, r := range key {
		// For each letter in the key,
		// find the child node of current node associated with that letter
		child := current.children[r]
		if child == nil {
			// If there is no child node associated with that letter,
			// create a new node and add it to current node as a child
			// associated with the letter
			child = NewTrieNode()
			current.children[r] = child
		}
		// Set current node = child node
		current = child
	}
	// Add value to current node
	current.entities.Add(entityID)

	// Release the exclusive lock
	trie.mx.Unlock()
	return nil
}

// findBranch is a helper function that traverses down
// to the node associated with a given key
func (trie *Trie) findBranch(key string) *TrieNode {
	// Read lock is not necessary, because findBranch will
	// only be called from functions with locks

	// Let current node = root node
	current := trie.root
	for _, r := range key {
		// Find the child node of the current node associated with that letter
		child := current.children[r]
		if child == nil {
			// If there is no child associated with that letter,
			// no keys start with the prefix, so return nil
			return nil
		}
		// Set current node = child node
		current = child
	}
	// Current now points to the branch containing all keys
	// that start with that prefix
	return current
}

// GetEntities returns the first n entity IDs that are associated with
// a given prefix string
// The actual work is done by the helper function getBranchEntities
func (trie *Trie) GetEntities(prefix string, limit int) []int {
	// Obtain a read lock
	trie.mx.RLock()
	// Use defer to exit the read lock as we exit the function
	defer trie.mx.RUnlock()

	prefixes := strings.Split(prefix, " ")

	// Child node now points to the branch containing all keys that
	// start with the prefix
	// recurse down the branch, gathering values
	branch := trie.findBranch(prefixes[0])
	if branch == nil {
		return []int{}
	}

	// Start with the ids that are returned for the first prefix.
	// We are guaranteed that our final result can only contain
	// less, and not more IDs.
	ei := NewEntityIdentifiers()
	getBranchEntities(branch, ei, limit)

	for i := 1; i < len(prefixes); i++ {
		branch := trie.findBranch(prefixes[i])
		if branch == nil {
			return []int{}
		}
		cei := NewEntityIdentifiers()
		getBranchEntities(branch, cei, limit)
		ei = ei.RetainAll(cei)
	}

	return ei.Sort()
}

// getBranchEntities is a helper function that returns the first n entity IDs that
// are associated with a given prefix string
func getBranchEntities(root *TrieNode, ei *EntityIdentifiers, limit int) {
	// Read lock is not necessary, because getBranchEntities will
	// only be called from functions with locks

	// We found entities, so we add them to our set
	// for as long as our limit permits
	if root.entities.Size() != 0 {
		cei := root.entities.Sort()
		for _, entityID := range cei {
			if limit == 0 || ei.Size() < limit {
				ei.Add(entityID)
			} else {
				return
			}
		}
	}

	// The current node has children, so we should traverse
	// down to them (and check if they have entities)
	if len(root.children) != 0 {
		keys := make([]rune, len(root.children))
		index := 0
		for k := range root.children {
			keys[index] = k
			index++
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, child := range keys {
			getBranchEntities(root.children[child], ei, limit)
		}
	}
}

// RemoveEntity removes a key/value pair from the trie
// The helper function removeEntity does the actual work
func (trie *Trie) RemoveEntity(key string, entityID int) error {
	if trie.root == nil || len(key) == 0 {
		return fmt.Errorf("Please provide a root and key")
	}
	trie.mx.Lock()
	defer trie.mx.Unlock()
	if err := removeEntity(trie.root, key, entityID); err != nil {
		return fmt.Errorf("Error removing entity: %v", err)
	}
	return nil
}

// removeEntity is a helper function that removes a key/value pair from the trie
func removeEntity(root *TrieNode, prefix string, entityID int) error {
	if prefix != "" {
		r, l := utf8.DecodeRuneInString(prefix)
		child := root.children[r]
		if child == nil {
			return fmt.Errorf("Key is not present in trie")
		}
		removeEntity(child, prefix[l:], entityID)
		if child.entities.Size() == 0 && len(child.children) == 0 {
			delete(root.children, r)
		}
	} else {
		root.entities.Remove(entityID)
	}
	return nil
}

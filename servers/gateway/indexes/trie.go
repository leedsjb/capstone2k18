package indexes

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

//Trie is an index used to look up the set of users that have
//a certain email/username/first name/last name
type Trie struct {
	root *TrieNode
	mx   sync.RWMutex
}

//NewTrie makes a new trie
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

//AddUser adds a user (ID) to the trie under a given key
func (trie *Trie) AddUser(key string, userID bson.ObjectId) error {
	if len(key) == 0 || len(userID) == 0 {
		return fmt.Errorf("Please provide a key and user ID")
	}

	//obtain an exclusive lock
	trie.mx.Lock()
	//let current node = root node
	current := trie.root

	for _, r := range key {
		//for each letter in the key
		//find the child node of current node associated with that letter
		child := current.children[r]
		if child == nil {
			//if there is no child node associated with that letter,
			//create a new node and add it to current node as a child
			//associated with the letter
			child = NewTrieNode()
			current.children[r] = child
		}
		//set current node = child node
		current = child
	}
	//add value to current node
	current.users.Add(userID)

	//release the exclusive lock
	trie.mx.Unlock()
	return nil
}

//findBranch is a helper function that traverses down
//to the node associated with a given key
func (trie *Trie) findBranch(key string) *TrieNode {
	//Read lock is not necessary, because findBranch will
	//only be called from functions with locks

	//let current node = root node
	current := trie.root
	for _, r := range key {
		//find the child node of the current node associated with that letter
		child := current.children[r]
		if child == nil {
			//if there is no child associated with that lette
			//no keys start with the prefix, so return nil
			return nil
		}
		//set current node = child node
		current = child
	}
	//current now points to the branch containing all keys
	//that start with that prefix
	return current
}

//GetUsers returns the first n user IDs that are associated with
//a given prefix string
//The actual work is done by the helper function getBranchUsers
func (trie *Trie) GetUsers(prefix string, limit int) []bson.ObjectId {
	//obtain a read lock
	trie.mx.RLock()
	//use defer to exit the read lock as we exit the function
	defer trie.mx.RUnlock()

	prefixes := strings.Split(prefix, " ")

	//child node now points to the branch containing all keys that
	//start with the prefix
	//recurse down the branch, gathering values
	branch := trie.findBranch(prefixes[0])
	if branch == nil {
		return []bson.ObjectId{}
	}

	//Start with the ids that are returned for the first prefix.
	//We are guaranteed that our final result can only contain
	//less, and not more IDs.
	ui := NewUserIdentifiers()
	getBranchUsers(branch, ui, limit)

	for i := 1; i < len(prefixes); i++ {
		branch := trie.findBranch(prefixes[i])
		if branch == nil {
			return []bson.ObjectId{}
		}
		cui := NewUserIdentifiers()
		getBranchUsers(branch, cui, limit)
		ui = ui.RetainAll(cui)
	}

	return ui.Sort()
}

//getBranchUsers is a helper function that returns the first n user IDs that
//are associated with a given prefix string
func getBranchUsers(root *TrieNode, ui *UserIdentifiers, limit int) {
	//Read lock is not necessary, because getBranchUsers will
	//only be called from functions with locks

	//We found users, so we add them to our set
	//for as long as our limit permits
	if root.users.Size() != 0 {
		cui := root.users.Sort()
		for _, userID := range cui {
			if limit == 0 || ui.Size() < limit {
				ui.Add(userID)
			} else {
				return
			}
		}
	}

	//The current node has children, so we should traverse
	//down to them (and check if they have users)
	if len(root.children) != 0 {
		keys := make([]rune, len(root.children))
		index := 0
		for k := range root.children {
			keys[index] = k
			index++
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, child := range keys {
			getBranchUsers(root.children[child], ui, limit)
		}
	}
}

//RemoveUser removes a key/value pair from the trie
//The helper function removeUser does the actual work
func (trie *Trie) RemoveUser(key string, userID bson.ObjectId) error {
	if trie.root == nil || len(key) == 0 || len(userID) == 0 {
		return fmt.Errorf("Please provide a root, key and user ID")
	}
	trie.mx.Lock()
	defer trie.mx.Unlock()
	if err := removeUser(trie.root, key, userID); err != nil {
		return fmt.Errorf("Error removing user: %v", err)
	}
	return nil
}

//removeUser is a helper function that removes a key/value pair from the trie
func removeUser(root *TrieNode, prefix string, userID bson.ObjectId) error {
	if prefix != "" {
		r, l := utf8.DecodeRuneInString(prefix)
		child := root.children[r]
		if child == nil {
			return fmt.Errorf("Key is not present in trie")
		}
		removeUser(child, prefix[l:], userID)
		if child.users.Size() == 0 && len(child.children) == 0 {
			delete(root.children, r)
		}
	} else {
		root.users.Remove(userID)
	}
	return nil
}

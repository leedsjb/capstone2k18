package indexes

import (
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

//TestNewTrie tests Trie's NewTrie function
func TestNewTrie(t *testing.T) {
	//Test if a struct, which we'll assume to be a Trie, is created
	trie := NewTrie()
	if trie == nil {
		t.Error("NewTrie() does not return a Trie")
	}
	//Test whether the root is initialized
	if trie.root == nil {
		t.Error("The trie is missing a root")
	}
}

//TestAddUser tests Trie's AddUser function
func TestAddUser(t *testing.T) {
	noUsers := []bson.ObjectId{}

	trie := NewTrie()

	//Test if the trie is empty at the start of the test
	if len(trie.root.children) != 0 {
		t.Error("The trie does not start out empty")
	}

	//Test if the add function throws an error if no key is provided
	if err := trie.AddUser("", bson.NewObjectId()); err == nil {
		t.Error("Didn't provide a key. Expected an error, but didn't get one")
	}

	//Test if the add function throws an error if no ID is provided
	if err := trie.AddUser("joey", ""); err == nil {
		t.Error("Didn't provide a ID. Expected an error, but didn't get one")
	}

	//Add the first user
	firstName1 := "matt"
	userID1 := bson.NewObjectId()
	trie.AddUser(firstName1, userID1)
	users1 := []bson.ObjectId{userID1}
	if err := trieContains(trie.root, firstName1, users1, false); err != nil {
		t.Errorf("User %v, %v was not added correctly: %v", firstName1, userID1, err)
	}

	//Add the second user
	firstName2 := "matthew"
	userID2 := bson.NewObjectId()
	trie.AddUser(firstName2, userID2)
	users2 := []bson.ObjectId{userID2}
	if err := trieContains(trie.root, firstName2, users2, false); err != nil {
		t.Errorf("User %v, %v was not added correctly: %v", firstName2, userID2, err)
	}

	//Test if "matt" (the first key) now indeed is a prefix
	if err := trieContains(trie.root, firstName1, users1, true); err != nil {
		t.Errorf("The trie was not built correctly: %v", err)
	}

	//Test if "matthe" (noninserted prefix) is now an existing prefix
	if err := trieContains(trie.root, "matthe", noUsers, true); err != nil {
		t.Errorf("The trie was not built correctly: %v", err)
	}

	//Add the third user
	firstName3 := "christine"
	userID3 := bson.NewObjectId()
	trie.AddUser(firstName3, userID3)
	users3 := []bson.ObjectId{userID3}
	if err := trieContains(trie.root, firstName3, users3, false); err != nil {
		t.Errorf("User %v, %v was not added correctly: %v", firstName3, userID3, err)
	}

	//Add the fourth user
	trie.AddUser(firstName3, userID3)
	if err := trieContains(trie.root, firstName3, users3, false); err != nil {
		t.Errorf("User %v, %v was not added correctly: %v", firstName3, userID3, err)
	}

	//Test if add handles identical keys, but different values, correctly
	firstName4 := firstName3
	userID4 := bson.NewObjectId()
	trie.AddUser(firstName4, userID4)
	users4 := []bson.ObjectId{userID3, userID4}
	if err := trieContains(trie.root, firstName4, users4, false); err != nil {
		t.Errorf("User %v, %v was not added correctly: %v", firstName4, userID4, err)
	}
}

//Test Trie's findBranch function
func TestFindBranch(t *testing.T) {
	trie := NewTrie()

	//Verify that a key of "" leads to the root
	if node := trie.findBranch(""); node != trie.root {
		t.Errorf("Expected: %v, but got: %v", trie.root, node)
	}

	//Add the first user
	firstName1 := "julie"
	userID1 := bson.NewObjectId()
	trie.AddUser(firstName1, userID1)
	expNode, _ := getNode(trie.root, "julie")
	if node := trie.findBranch("julie"); node != expNode {
		t.Errorf("Expected: %v, but got: %v", expNode, node)
	}

	//Add the second user
	firstName2 := "juliet"
	userID2 := bson.NewObjectId()
	trie.AddUser(firstName2, userID2)
	expNode, _ = getNode(trie.root, "julie")
	if node := trie.findBranch("julie"); node != expNode {
		t.Errorf("Expected: %v, but got: %v", expNode, node)
	}
}

//Test Trie's getBranchUsers function
func TestGetBranchUsers(t *testing.T) {
	trie := NewTrie()

	//Create a new set
	ui := NewUserIdentifiers()
	getBranchUsers(trie.root, ui, 0)
	//Test if an empty trie indeed leads to a result of 0 users
	if !reflect.DeepEqual(ui.Sort(), []bson.ObjectId{}) {
		t.Errorf("Expected: %v, but got: %v", []bson.ObjectId{}, ui.Sort())
	}

	//Create a new set
	ui = NewUserIdentifiers()
	getBranchUsers(trie.root, ui, 1)

	//Test if a limit of 1 still leads to a result of 0 users
	//for an empty trie
	if !reflect.DeepEqual(ui.Sort(), []bson.ObjectId{}) {
		t.Errorf("Expected: %v, but got: %v", []bson.ObjectId{}, ui.Sort())
	}

	//Add the first user
	firstName1 := "jon"
	userID1 := bson.NewObjectId()
	trie.AddUser(firstName1, userID1)

	//Add the second user
	firstName2 := firstName1
	userID2 := bson.NewObjectId()
	trie.AddUser(firstName2, userID2)

	//Add the third user. This is a duplicate
	firstName3 := firstName1
	userID3 := userID1
	trie.AddUser(firstName3, userID3)

	//Add the fourth user
	firstName4 := "jonathan"
	userID4 := bson.NewObjectId()
	trie.AddUser(firstName4, userID4)

	//Add the fifth user
	firstName5 := "bryn"
	userID5 := bson.NewObjectId()
	trie.AddUser(firstName5, userID5)

	//Add the sixth user
	firstName6 := "brian"
	userID6 := bson.NewObjectId()
	trie.AddUser(firstName6, userID6)

	//Set up the slice of expected users
	expUsers1 := []bson.ObjectId{userID1, userID2, userID4, userID5, userID6}

	//Create a new slice
	ui = NewUserIdentifiers()
	getBranchUsers(trie.root, ui, 0)

	if !reflect.DeepEqual(ui.Sort(), expUsers1) {
		t.Errorf("Expected: %v, but got: %v", expUsers1, ui.Sort())
	}

	expUsers2 := []bson.ObjectId{userID6}

	ui = NewUserIdentifiers()
	getBranchUsers(trie.root, ui, 1)
	if !reflect.DeepEqual(expUsers2, ui.Sort()) {
		t.Errorf("Expected: %v, but got: %v", expUsers2, ui.Sort())
	}

	//Test if a leaf with a single user only returns one user
	nodeUser4, _ := getNode(trie.root, firstName4)
	expUsers3 := []bson.ObjectId{userID4}
	ui = NewUserIdentifiers()
	getBranchUsers(nodeUser4, ui, 0)
	if !reflect.DeepEqual(expUsers3, ui.Sort()) {
		t.Errorf("Expected: %v, but got: %v", expUsers3, ui.Sort())
	}

	//See if the correct users are returned for an "arbitrary" prefix
	prefixNode, _ := getNode(trie.root, "br")
	expUsers4 := []bson.ObjectId{userID5, userID6}
	ui = NewUserIdentifiers()
	getBranchUsers(prefixNode, ui, 0)
	if !reflect.DeepEqual(expUsers4, ui.Sort()) {
		t.Errorf("Expected: %v, but got: %v", expUsers4, ui.Sort())
	}
}

//TestGetUsers tests Trie's GetUsers function
func TestGetUsers(t *testing.T) {
	//Only test what happens for a non existing string
	trie := NewTrie()

	firstName1 := "jonathan"
	userID1 := bson.NewObjectId()
	trie.AddUser(firstName1, userID1)

	firstname2 := "jon"
	userID2 := bson.NewObjectId()
	trie.AddUser(firstname2, userID2)

	firstName3 := "jonny"
	userID3 := bson.NewObjectId()
	trie.AddUser(firstName3, userID3)

	lastName1 := "mack"
	trie.AddUser(lastName1, userID3)

	//Test if looking for a non-existing key indeed returns 0 users
	if users := trie.GetUsers("eli", 0); !reflect.DeepEqual([]bson.ObjectId{}, users) {
		t.Errorf("Expected: %v, but got: %v", []bson.ObjectId{}, users)
	}

	//Test if the correct user is returned when someone looks for firstName + lastName
	//This tests if we implemented the first extra credit correctly
	//Simultaneously check if casing matters
	expUsers1 := []bson.ObjectId{userID3}
	if users := trie.GetUsers("jonny mack", 0); !reflect.DeepEqual(expUsers1, users) {
		t.Errorf("Expected: %v, but got: %v", expUsers1, users)
	}
}

//TestRemoveUsers tests the RemoveUsers function
func TestRemoveUsers(t *testing.T) {
	//Create a new trie
	trie := NewTrie()

	//Start adding a wide variety of users
	firstName1 := "brady"
	userID1 := bson.NewObjectId()
	trie.AddUser(firstName1, userID1)

	firstName2 := "dann"
	userID2 := bson.NewObjectId()
	trie.AddUser(firstName2, userID2)

	firstName3 := "dani"
	userID3 := bson.NewObjectId()
	trie.AddUser(firstName3, userID3)

	firstName4 := "katie"
	userID4 := bson.NewObjectId()
	trie.AddUser(firstName4, userID4)

	firstName5 := "dylan"
	userID5 := bson.NewObjectId()
	trie.AddUser(firstName5, userID5)

	firstName6 := "justin"
	userID6 := bson.NewObjectId()
	trie.AddUser(firstName6, userID6)

	firstName7 := firstName6
	userID7 := bson.NewObjectId()
	trie.AddUser(firstName7, userID7)

	firstName8 := "max"
	userID8 := bson.NewObjectId()
	trie.AddUser(firstName8, userID8)

	firstName9 := firstName8
	userID9 := bson.NewObjectId()
	trie.AddUser(firstName9, userID9)

	firstName10 := "maximilian"
	userID10 := bson.NewObjectId()
	trie.AddUser(firstName10, userID10)

	firstName11 := "daniel"
	userID11 := bson.NewObjectId()
	trie.AddUser(firstName11, userID11)

	firstName12 := "brady"
	userID12 := bson.NewObjectId()
	trie.AddUser(firstName12, userID12)

	//Test if an error is returned for the absence of a key
	if err := trie.RemoveUser("", userID1); err == nil {
		t.Error("Attempted to remove a user without a key. Expected an error but didn't get one")
	}

	//Test if an error is returned for the absence of a user ID
	if err := trie.RemoveUser(firstName1, ""); err == nil {
		t.Error("Attempted to remove a user without a ID. Expected an error but didn't get one")
	}

	//Get the trie's original satte
	ogState := trie.GetUsers("", 0)

	//Try to remove a non-existing key
	newState1 := trie.GetUsers("", 0)
	if !reflect.DeepEqual(ogState, newState1) {
		t.Errorf("Expected: %v and %v", ogState, newState1)
	}

	//Trie to remove a key that is only a prefix
	newState2 := trie.GetUsers("", 0)
	if !reflect.DeepEqual(ogState, newState2) {
		t.Errorf("Expected: %v and %v", ogState, newState2)
	}

	//Remove a leaf with only one user
	expState1 := []bson.ObjectId{userID1, userID2, userID3, userID5, userID6, userID7, userID8, userID9, userID10, userID11, userID12}
	if trie.RemoveUser(firstName4, userID4); !reflect.DeepEqual(expState1, trie.GetUsers("", 0)) {
		t.Errorf("Expected: %v, but got: %v", expState1, trie.GetUsers("", 0))
	}

	//Test if the trie cleans up after itself
	if node, path := getNode(trie.root, firstName4[:len(firstName4)-1]); node != nil || path == firstName4[:len(firstName4)-1] {
		t.Errorf("Expected: %v, but got: %v. The traversed path was: %v, which should not have existed.", nil, node, path)
	}

	//Remove a key that sits at a leaf with more than one user
	expState2 := []bson.ObjectId{userID1, userID2, userID3, userID5, userID7, userID8, userID9, userID10, userID11, userID12}
	if trie.RemoveUser(firstName6, userID6); !reflect.DeepEqual(expState2, trie.GetUsers("", 0)) {
		t.Errorf("Expected: %v, but got: %v", expState2, trie.GetUsers("", 0))
	}

	//Remove a key that is a prefix with only one user
	expState3 := []bson.ObjectId{userID1, userID2, userID5, userID7, userID8, userID9, userID10, userID11, userID12}
	if trie.RemoveUser(firstName3, userID3); !reflect.DeepEqual(expState3, trie.GetUsers("", 0)) {
		t.Errorf("Expected: %v, but got: %v", expState3, trie.GetUsers("", 0))
	}

	//Remove a key that is a prefix with more than one user
	expState4 := []bson.ObjectId{userID1, userID2, userID5, userID7, userID9, userID10, userID11, userID12}
	if trie.RemoveUser(firstName8, userID8); !reflect.DeepEqual(expState4, trie.GetUsers("", 0)) {
		t.Errorf("Expected: %v, but got: %v", expState4, trie.GetUsers("", 0))
	}

	//Destination node has more than one user, but not this user
	ogState = trie.GetUsers("", 0)
	trie.RemoveUser(firstName12, bson.NewObjectId())
	newState3 := trie.GetUsers("", 0)
	if !reflect.DeepEqual(ogState, newState3) {
		t.Errorf("Expected: %v and %v", ogState, newState3)
	}

	//Start emptying the trie
	expState5 := []bson.ObjectId{}
	trie.RemoveUser(firstName1, userID1)
	trie.RemoveUser(firstName2, userID2)
	trie.RemoveUser(firstName5, userID5)
	trie.RemoveUser(firstName7, userID7)
	trie.RemoveUser(firstName9, userID9)
	trie.RemoveUser(firstName10, userID10)
	trie.RemoveUser(firstName11, userID11)
	trie.RemoveUser(firstName12, userID12)

	//Test if the trie is indeed empty
	if !reflect.DeepEqual(trie.GetUsers("", 0), expState5) {
		t.Errorf("Expected: %v, but got: %v", expState5, trie.GetUsers("", 0))
	}

	//Test if the root is still present after an attempt to delete
	//a non existing user
	trie.RemoveUser("geoff", bson.NewObjectId())
	if trie.root == nil {
		t.Errorf("Trie's root got removed unexpectedly")
	}

	//Verify that the state is still empty
	if !reflect.DeepEqual(trie.GetUsers("", 0), expState5) {
		t.Errorf("Expected: %v, but got: %v", expState5, trie.GetUsers("", 0))
	}
}

//getNode is a helper function, specifically written for these tests, that
//returns a node for any given key
func getNode(root *TrieNode, key string) (*TrieNode, string) {
	path := ""

	current := root
	for _, r := range key {
		child := current.children[r]
		if child == nil {
			//Key does not exist in trie
			return nil, path
		}
		current = child
		path += string(r)
	}

	return current, path
}

//trieContains is a helper function, specifically written for these tests, that
//extensively tests if a node appears in the trie as we expect it to be
func trieContains(root *TrieNode, firstName string, users []bson.ObjectId, expectChildren bool) error {
	rUser, path := getNode(root, firstName)
	if rUser == nil {
		return fmt.Errorf("Trie doesn't contain user")
	}
	if path != firstName {
		return fmt.Errorf("Test did not check the correct location")
	}
	if !reflect.DeepEqual(rUser.users.Sort(), users) {
		return fmt.Errorf("Expected: %v, but got: %v", users, rUser.users.Sort())
	}
	if expectChildren && len(rUser.children) == 0 {
		return fmt.Errorf("Node is not a prefix, but should be")
	}
	if !expectChildren && len(rUser.children) != 0 {
		return fmt.Errorf("Node is a prefix, but should not be")
	}
	return nil
}

package users

import (
	"errors"

	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/indexes"

	"gopkg.in/mgo.v2/bson"
)

//ErrUserNotFound is returned when the user can't be found
var ErrUserNotFound = errors.New("user not found")

//Store represents a store for Users
type Store interface {
	//GetByID returns the User with the given ID
	GetByID(id bson.ObjectId) (*User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*User, error)

	//Insert converts the NewUser to a User, inserts
	//it into the database, and returns it
	Insert(newUser *NewUser) (*User, error)

	//Update applies UserUpdates to the given user ID
	Update(userID bson.ObjectId, updates *Updates) error

	//Delete deletes the user with the given ID
	Delete(userID bson.ObjectId) error

	//AddUserToTrie adds a user to the trie
	AddUserToTrie(trie *indexes.Trie, user *User) error

	//LoadTrie fills a trie based on existing user accounts
	LoadTrie(trie *indexes.Trie) error

	//GetUsers retrieves users from the store
	GetUsers(users []bson.ObjectId) (*[]*User, error)
}

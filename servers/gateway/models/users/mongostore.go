package users

import (
	"fmt"
	"strings"

	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/indexes"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoStore implements Store for MongoDB
type MongoStore struct {
	//mongo session
	session *mgo.Session
	//database name
	dbname string
	//collection name
	colname string
	//Collection object
	col *mgo.Collection
}

//NewMongoStore creates a new Mongo store
func NewMongoStore(sess *mgo.Session, dbName string, colName string) *MongoStore {
	return &MongoStore{
		session: sess,
		dbname:  dbName,
		colname: colName,
		col:     sess.DB(dbName).C(colName),
	}
}

//Insert converts the NewUser to a User, inserts
//it into the database, and returns it
func (s *MongoStore) Insert(newUser *NewUser) (*User, error) {
	u, err := newUser.ToUser()
	if err != nil {
		return nil, fmt.Errorf("Error converting new user: %v", err)
	}
	//insert the User into the database and return it,
	//or an error if one occurred
	if err := s.col.Insert(u); err != nil {
		return nil, fmt.Errorf("Error inserting user: %v", err)
	}
	return u, nil
}

//GetByID returns the User with the given ID
//Returns an error if the user could not be found
func (s *MongoStore) GetByID(id bson.ObjectId) (*User, error) {
	u := &User{}
	if err := s.col.FindId(id).One(u); err != nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}

//GetByField returns the first user for which the given field
//has the given found
//Returns an error if no user matches the query
func (s *MongoStore) GetByField(field string, value string) (*User, error) {
	u := &User{}
	if err := s.col.Find(bson.M{field: value}).One(u); err != nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}

//GetByEmail returns the user with the given email
//Returns an error if the user could not be found
func (s *MongoStore) GetByEmail(email string) (*User, error) {
	u, err := s.GetByField("email", email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//GetByUserName returns the User with the given Username
//Returns an error if the user could not be found
func (s *MongoStore) GetByUserName(username string) (*User, error) {
	u, err := s.GetByField("username", username)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Update applies UserUpdates to the given user ID
//Per updates, only a user's first name and/or last name can be
//updated
func (s *MongoStore) Update(userID bson.ObjectId, updates *Updates) error {
	if err := s.col.UpdateId(userID, bson.M{"$set": updates}); err != nil {
		return fmt.Errorf("Error updating user: %v", err)
	}
	return nil
}

//Delete deletes the user with the given ID
func (s *MongoStore) Delete(userID bson.ObjectId) error {
	if err := s.col.RemoveId(userID); err != nil {
		return fmt.Errorf("Error deleting user: %v", err)
	}
	return nil
}

//AddUserToTrie adds a mongo document (user) to the trie
func AddUserToTrie(trie *indexes.Trie, user *User) error {
	if err := trie.AddUser(strings.ToLower(user.Email), user.ID); err != nil {
		return fmt.Errorf("Error adding user to trie: %v", err)
	}
	if err := trie.AddUser(strings.ToLower(user.UserName), user.ID); err != nil {
		return fmt.Errorf("Error adding user to trie: %v", err)
	}
	if err := trie.AddUser(strings.ToLower(user.FirstName), user.ID); err != nil {
		return fmt.Errorf("Error adding user to trie: %v", err)
	}
	if err := trie.AddUser(strings.ToLower(user.LastName), user.ID); err != nil {
		return fmt.Errorf("Error adding user to trie: %v", err)
	}
	return nil
}

//LoadTrie generates a new trie based on existing user accounts
func (s *MongoStore) LoadTrie(trie *indexes.Trie) error {
	user := &User{}

	iter := s.col.Find(nil).Iter()
	for iter.Next(user) {
		if err := AddUserToTrie(trie, user); err != nil {
			return fmt.Errorf("Error loading trie: %v", err)
		}
	}
	if err := iter.Close(); err != nil {
		return err
	}

	return nil
}

//GetUsers retrieves users from the database for a list of IDs
func (s *MongoStore) GetUsers(userIDS []bson.ObjectId) ([]*User, error) {
	var results = []*User{}

	for _, userID := range userIDS {
		user, err := s.GetByID(userID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving user from store: %v", err)
		}
		results = append(results, user)
	}

	return results, nil
}

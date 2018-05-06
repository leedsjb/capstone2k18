package users

import "gopkg.in/mgo.v2/bson"
import "fmt"

//MemStore is a local user store
type MemStore struct {
	collection []*User
}

//Insert converts the NewUser to a User, inserts
//it into the database, and returns it
func (s *MemStore) Insert(newUser *NewUser) (*User, error) {
	u, err := newUser.ToUser()
	if err != nil {
		return nil, fmt.Errorf("Error converting user: %v", err)
	}
	s.collection = append(s.collection, u)
	//Return the just inserted user
	return s.collection[len(s.collection)-1], nil
}

//GetByID returns the User with the given ID
//Returns an error if the user could not be found
func (s *MemStore) GetByID(id bson.ObjectId) (*User, error) {
	for _, u := range s.collection {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

//GetByEmail returns the user with the given email
//Returns an error if the user could not be found
func (s *MemStore) GetByEmail(email string) (*User, error) {
	for _, u := range s.collection {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

//GetByUserName returns the User with the given Username
//Returns an error if the user could not be found
func (s *MemStore) GetByUserName(username string) (*User, error) {
	for _, u := range s.collection {
		if u.UserName == username {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

//Update applies UserUpdates to the given user ID
//Per updates, only a user's first name and/or last name can be
//updated
func (s *MemStore) Update(userID bson.ObjectId, updates *Updates) error {
	u, err := s.GetByID(userID)
	if err != nil {
		return fmt.Errorf("Error retrieving ID: %v", err)
	}
	if err := u.ApplyUpdates(updates); err != nil {
		return fmt.Errorf("Error applying updates: %v", err)
	}
	return nil
}

//Delete deletes the user with the given ID
func (s *MemStore) Delete(userID bson.ObjectId) error {
	index := -1
	for i, u := range s.collection {
		if u.ID == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return ErrUserNotFound
	}
	//https://github.com/golang/go/wiki/SliceTricks
	//Delete user without leaving gaps in data structure
	s.collection = append(s.collection[:index], s.collection[index+1:]...)
	return nil
}

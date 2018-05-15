package users

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

//Store information about the store
var s *MongoStore

//TestMain sets up the Mongo store so that other functions don't have to
func TestMain(m *testing.M) {
	//get the address of the MongoDB server
	//from an environment variable
	memStore = &MemStore{}
	mongoAddr := os.Getenv("MONGO_ADDR")
	//default to "localhost"
	if len(mongoAddr) == 0 {
		mongoAddr = "localhost"
	}
	mongoSess, err := mgo.Dial(mongoAddr)
	if err != nil {
		fmt.Printf("Error dialing mongo: %v", err)
	}
	s = NewMongoStore(mongoSess, "users", "users")
	os.Exit(m.Run())
}

//TestInsert tests the insert function
func TestInsert(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	if _, err := s.Insert(nu); err != nil {
		t.Errorf("Error inserting user: %v", err)
	}
	//Remove inserted documents
	s.col.RemoveAll(nil)
}

//TestGetByID tests the GetByID function
//Checks if the function returns the correct user for a present ID and throws
//an error when used with a non-present ID
func TestGetByID(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := s.Insert(nu)
	if err != nil {
		t.Errorf("Error inserting user: %v", err)
	}
	//Ensure we have the ID of the just inserted user
	rid := u.ID

	cases := []struct {
		name        string
		id          bson.ObjectId
		expectError bool
	}{
		{
			"present ID",
			rid,
			false,
		},
		{
			"non-present ID",
			bson.NewObjectId(),
			true,
		},
	}

	//Go over all the test cases
	for _, c := range cases {
		ru, err := s.GetByID(c.id)
		if err != nil {
			if !c.expectError {
				t.Errorf("Unexpected error while retrieving user")
			}
		} else if !reflect.DeepEqual(u, ru) {
			t.Error("Retrieved user does not match inserted user")
		}
	}
	//Remove inserted documents
	s.col.RemoveAll(nil)
}

//TestGetByEmail tests the GetByEmail function, which is used to retrieve
//a user by their email address
//Verifies that it works for a present email and throws an error for a
//non-present email
func TestGetByEmail(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := s.Insert(nu)
	if err != nil {
		t.Errorf("Error inserting user: %v", err)
	}
	rEmail := u.Email
	cases := []struct {
		name        string
		email       string
		expectError bool
	}{
		{
			"present email",
			rEmail,
			false,
		},
		{
			"non-present email",
			"andrewwiles@gmail.com",
			true,
		},
	}

	//Go over all of the test cases
	for _, c := range cases {
		ru, err := s.GetByEmail(c.email)
		if err != nil {
			if !c.expectError {
				t.Errorf("Unexpected error while retrieving user by email: %v", err)
			}
		} else {
			//Check if the correct user was retrieved
			if !reflect.DeepEqual(u, ru) {
				t.Error("Retrieved user does not match inserted user")
			}
		}
	}
	//Remove inserted documents
	s.col.RemoveAll(nil)
}

//TestGetByUsername tests the TestGetByUsername function, that is responsible
//for retrieving a user by their username
//Verifies that the function works for a present username and throws an error
//when there is not a user present with the given username
func TestGetByUserName(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := s.Insert(nu)
	if err != nil {
		t.Errorf("Error inserting user: %v", err)
	}
	//Obtain the username of the inserted user
	rUsername := u.UserName
	cases := []struct {
		name        string
		username    string
		expectError bool
	}{
		{
			"present username",
			rUsername,
			false,
		},
		{
			"non-present username",
			"Leon",
			true,
		},
	}

	//Go over the test cases
	for _, c := range cases {
		ru, err := s.GetByUserName(c.username)
		if err != nil {
			if !c.expectError {
				t.Errorf("Unexpected error while retrieving user by username: %v", err)
			}
		} else {
			//Ensure that correct user was retrieved
			if !reflect.DeepEqual(u, ru) {
				t.Error("Retrieved user does not match inserted user")
			}
		}
	}

	//Remove inserted documents
	s.col.RemoveAll(nil)
}

//TestUpdate tests the Update function, which is used to update a user's first name and/or
//last name
//Verifies that the function works if, and only if, both a (new) first name and
//last name are present
func TestUpdate(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := s.Insert(nu)
	if err != nil {
		t.Errorf("Error inserting user: %v", err)
	}
	//Retrieve the ID of the just inserted user
	rid := u.ID

	cases := []struct {
		name         string
		id           bson.ObjectId
		newFirstName string
		newLastname  string
		expectError  bool
	}{
		{
			"valid update, incorrect ID",
			bson.NewObjectId(),
			"Pierre",
			"de Fermat",
			true,
		},
		{
			"valid update",
			rid,
			"Pierre",
			"de Fermat",
			false,
		},
		{
			"invalid update, empty first name",
			rid,
			"",
			"de Fermat",
			true,
		},
		{
			"invalid update, empty last name",
			rid,
			"Pierre",
			"",
			true,
		},
	}

	//Go over test cases
	for _, c := range cases {
		//Ensure initial first name and last name are used
		//Necesary to ensure test independency
		if u.FirstName != firstName || u.LastName != lastName {
			t.Errorf("Update was not reset in test: %v", err)
		}
		updates := &Updates{FirstName: c.newFirstName, LastName: c.newLastname}
		if err := s.Update(c.id, updates); err != nil {
			if !c.expectError {
				t.Errorf("Error updating user: %v", err)
			}
			//We failed to update a user profile, so we can quit the test
			continue
		}
		//Retrieve the user profile that was updated
		ru, err := s.GetByID(c.id)
		if err != nil {
			if !c.expectError {
				t.Errorf("Unexpected error while retrieving user: %v", err)
			}
			//We failed to retrieve the updated user's profile, so we can
			//quit the test
			continue
		}
		//Verify that the first name was updated correctly
		if ru.FirstName != c.newFirstName {
			t.Error("First name was not updated correctly")
		}
		//Verify that the last name was updated correctly
		if ru.LastName != c.newLastname {
			t.Error("Last name was not updated correctly")
		}
		//Reset user to original state
		//Necessary for test independency
		updates = &Updates{FirstName: firstName, LastName: lastName}
		if err := s.Update(c.id, updates); err != nil {
			t.Errorf("Error updating user to original state: %v", err)
		}
	}
	//Remove the inserted documents
	s.col.RemoveAll(nil)
}

//TestDelete tests the delete function
//Verifies that a user profile is correctly deleted if, and only if,
//a present ID is supplied
func TestDelete(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := s.Insert(nu)
	if err != nil {
		t.Errorf("Error inserting user: %v", err)
	}
	id := u.ID

	cases := []struct {
		name        string
		id          bson.ObjectId
		expectError bool
	}{
		{
			"valid ID",
			id,
			false,
		},
		{
			"invalid ID",
			bson.NewObjectId(),
			true,
		},
	}

	//Go over test cases
	for _, c := range cases {
		err = s.Delete(id)

		if err != nil {
			if !c.expectError {
				t.Errorf("Unexpected error while deleting user: %v", err)
			}
			//We failed to delete a user, so there is no point in continuing
			//our tests
			continue
		}

		//We successfully deleted a user, so a Get should now return an error
		if _, err := s.GetByID(c.id); err == nil {
			t.Errorf("Expected an error while retrieving a deleted user")
		}
	}
	//Remove just inserted documents
	s.col.RemoveAll(nil)
}

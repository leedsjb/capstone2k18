package users

import (
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

var memStore *MemStore

//TestMemInsert tests the insert function
func TestMemInsert(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	if _, err := memStore.Insert(nu); err != nil {
		t.Errorf("Error inserting user: %v", err)
	}
	//Remove inserted user
	memStore.collection = nil
}

//TestMemGetByID tests the TestGetByID function
//Checks if the function returns the correct user for a present ID and throws
//an error when used with a non-present ID
func TestMemGetByID(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := memStore.Insert(nu)
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
			"present id",
			rid,
			false,
		},
		{
			"non-present id",
			bson.NewObjectId(),
			true,
		},
	}

	//Go over all the test cases
	for _, c := range cases {
		ru, err := memStore.GetByID(c.id)
		if err != nil {
			if !c.expectError {
				t.Errorf("Error retrieving user: %v", err)
			}
			continue
		}
		if !reflect.DeepEqual(u, ru) {
			t.Error("Incorrect user was retrieved")
		}
	}
	//Remove inserted user
	memStore.collection = nil
}

//TestMemGetByEmail tests the GetByEmail function, which is used to retrieve
//a user by their email address
//Verifies that it works for a present email and throws an error for a
//non-present email
func TestMemGetByEmail(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := memStore.Insert(nu)
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

	//Go over test cases
	for _, c := range cases {
		ru, err := memStore.GetByEmail(c.email)
		if err != nil {
			if !c.expectError {
				t.Errorf("Error retrieving user: %v", err)
			}
			continue
		}
		//Check if the correct user was retrieved
		if !reflect.DeepEqual(u, ru) {
			t.Error("Incorrect user was retrieved")
		}
	}
	//Remove inserted user
	memStore.collection = nil
}

//TestMemGetByUsername tests the TestGetByUsername function, that is responsible
//for retrieving a user by their username
//Verifies that the function works for a present username and throws an error
//when there is not a user present with the given username
func TestMemGetByUserName(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := memStore.Insert(nu)
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

	//Go over test cases
	for _, c := range cases {
		ru, err := memStore.GetByUserName(c.username)
		if err != nil {
			if !c.expectError {
				t.Errorf("Error retrieving user: %v", err)
			}
			continue
		}
		if !reflect.DeepEqual(u, ru) {
			t.Error("Incorrect user was retrieved")
		}
	}

	//Remove the inserted user
	memStore.collection = nil
}

//TestMemUpdate tests the Update function, which is used to update a user's
//first name and/or last name
//Verifies that the function works if, and only if, both a (new) first name and
//last name are present
func TestMemUpdate(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := memStore.Insert(nu)
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
		if err := memStore.Update(c.id, updates); err != nil {
			if !c.expectError {
				t.Errorf("Error updating user: %v", err)
			}
			//We failed to update a user profile, so we can quit the test
			continue
		}
		//Retrieve the user profile that was updated
		ru, err := memStore.GetByID(c.id)
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
		if err := memStore.Update(c.id, updates); err != nil {
			t.Errorf("Error updating user to original state: %v", err)
		}
	}
	memStore.collection = nil
}

//TestDelete tests the delete function
//Verifies that a user profile is correctly deleted if, and only if,
//a present ID is supplied
func TestMemDelete(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := memStore.Insert(nu)
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
		err = memStore.Delete(id)

		if err != nil {
			if !c.expectError {
				t.Errorf("Unexpected error while deleting user: %v", err)
			}
			//We failed to delete a user, so there is no point in continuing
			//our tests
			continue
		}

		//We successfully deleted a user, so a Get should now return an error
		if _, err := memStore.GetByID(c.id); err == nil {
			t.Errorf("Expected an error while retrieving a deleted user")
		}
	}

	//Remove the just inserted user
	memStore.collection = nil
}

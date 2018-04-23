package users

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

//add tests for the various functions in user.go, as described in the assignment.
//use `go test -cover` to ensure that you are covering all or nearly all of your code paths.

//TestValidate tests Validate, the function that checks a user profile
//Checks if all the validation errors are returned where necessary
//And that no errors are returned for valid Users
func TestValidate(t *testing.T) {
	cases := []struct {
		name         string
		email        string
		password     string
		passwordConf string
		username     string
		expectError  bool
	}{
		{
			"no valid email address",
			emailInvalid,
			passValid,
			passValid,
			usernameValid,
			true,
		},
		{
			"password less than 6 characters",
			emailValid,
			passInvalid,
			passInvalid,
			usernameValid,
			true,
		},
		{
			"passwords do not match",
			emailValid,
			passValid,
			passValid2,
			usernameValid,
			true,
		},
		{
			"username is zero length",
			emailValid,
			passValid,
			passValid,
			usernameInvalid,
			true,
		},
		{
			"valid user",
			emailValid,
			passValid,
			passValid,
			usernameValid,
			false,
		},
	}
	for _, c := range cases {
		u := &NewUser{Email: c.email, Password: c.password,
			PasswordConf: c.passwordConf, UserName: c.username}
		err := u.Validate()
		if err != nil {
			if !c.expectError {
				t.Errorf("Unexpected error while validating user: %v", err)
			}
		} else if c.expectError {
			t.Errorf("Expected an error but didn't get one")
		}
	}
}

//TestToUser() tests the ToUser function, that converts a new user to a
//real user.
//Checks if PhotoURL and PassHash are calculated correctly.
func TestToUser(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid, PasswordConf: passValid2,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	h := md5.New()
	h.Write([]byte(strings.ToLower(strings.TrimSpace(nu.Email))))
	hashExpected := gravatarBasePhotoURL + hex.EncodeToString(h.Sum(nil))
	u, err := nu.ToUser()
	if err != nil {
		t.Errorf("Error converting new user: %v", err)
	} else {
		if hashExpected != u.PhotoURL {
			t.Errorf("The Gravatar photo url is not generated properly")
		}
		if err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(passValid)); err != nil {
			t.Errorf("Incorrect password hash was generated")
		}
	}
}

//TestFullName tests the full name function, that returns a user's full name.
//Checks if the full name format is correct
func TestFullName(t *testing.T) {
	cases := []struct {
		name           string
		firstName      string
		lastName       string
		expectedOutput string
	}{
		{
			"no first name",
			"",
			lastName,
			lastName,
		},
		{
			"no last name",
			firstName,
			"",
			firstName,
		},
		{
			"valid full name",
			firstName,
			lastName,
			firstName + " " + lastName,
		},
	}
	for _, c := range cases {
		u := &User{FirstName: c.firstName, LastName: c.lastName}
		if output := u.FullName(); output != c.expectedOutput {
			t.Errorf("Incorrect output for `%s` and `%s`: expected `%s` but got `%s`",
				c.firstName, c.lastName, c.expectedOutput, output)
		}
	}
}

//TestAuthenticate tests the authenticate function, that validates a provided
//password by crossreferencing a hash.
func TestAuthenticate(t *testing.T) {
	nu := &NewUser{Email: emailValid, Password: passValid2, PasswordConf: passValid2,
		UserName: usernameValid, FirstName: firstName, LastName: lastName}
	u, err := nu.ToUser()
	if err != nil {
		t.Errorf("Error converting new user: %v", err)
	} else {
		if err := u.Authenticate(passValid); err == nil {
			t.Errorf("Incorrect password is not generating an authentication error")
		}
		if err := u.Authenticate(passValid2); err != nil {
			t.Errorf("Error authenticating the correct password")
		}
	}
}

//TestApplyUpdates tests the ApplyUpdates function, that is used to change a
//user's first name and/or last name
//Verifies that the absence of either a first name or last name throws an error
func TestApplyUpdates(t *testing.T) {
	cases := []struct {
		name              string
		expectedFirstName string
		expectedLastName  string
		expectError       bool
	}{
		{
			"valid update",
			"Pierre",
			"de Fermat",
			false,
		},
		{
			"empty first name",
			"",
			"de Fermat",
			true,
		},
		{
			"empty last name",
			"Pierre",
			"",
			true,
		},
	}

	for _, c := range cases {
		nu := &NewUser{Email: emailValid, Password: passValid,
			PasswordConf: passValid, UserName: usernameValid, LastName: lastName}
		u, err := nu.ToUser()
		if err != nil {
			t.Errorf("Error converting new user: %v", err)
		} else {
			updates := &Updates{FirstName: c.expectedFirstName,
				LastName: c.expectedLastName}
			err := u.ApplyUpdates(updates)
			if err != nil {
				if !c.expectError {
					t.Errorf("Unexpected error while applying updates: %v", err)
				}
				continue
			} else if c.expectError {
				t.Errorf("Didn't get an error but expected one")
			}
			if u.FirstName != c.expectedFirstName || u.LastName != c.expectedLastName {
				t.Errorf("The updates did not get applied correctly")
			}
		}
	}
}

package users

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

var bcryptCost = 13

//User represents a user account in the database
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email"`
	PassHash  []byte        `json:"-"` //stored, but not encoded to clients
	UserName  string        `json:"userName"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	PhotoURL  string        `json:"photoURL"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//Updates represents allowed updates to a user profile
type Updates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//Validate validates the new user and returns an error if
//any of the validation rules fail, or nil if its valid
func (nu *NewUser) Validate() error {
	//validate the new user according to these rules:
	//use fmt.Errorf() to generate appropriate error messages if
	//the new user doesn't pass one of the validation rules
	//Email field must be a valid email address
	if _, err := mail.ParseAddress(nu.Email); err != nil {
		return fmt.Errorf("Email address must be valid: %v", err)
	}
	//Password must be at least 6 characters
	if len(nu.Password) < 6 {
		return fmt.Errorf("Password must be at least 6 characters")
	}
	//Password and PasswordConf must match
	if nu.Password != nu.PasswordConf {
		return fmt.Errorf("Passwords must match")
	}
	//UserName must be non-zero length
	if len(nu.UserName) == 0 {
		return fmt.Errorf("Username must be non-zero length")
	}
	return nil
}

//ToUser converts the NewUser to a User, setting the
//PhotoURL and PassHash fields appropriately
func (nu *NewUser) ToUser() (*User, error) {
	//set the PhotoURL field of the new User to
	//the Gravatar PhotoURL for the user's email address.
	h := md5.New()
	if _, err := h.Write([]byte(strings.ToLower(strings.TrimSpace(nu.Email)))); err != nil {
		return nil, fmt.Errorf("Error hashing email: %v", err)
	}
	pu := gravatarBasePhotoURL + hex.EncodeToString(h.Sum(nil))
	u := &User{
		//set the ID field of the new User
		//to a new bson ObjectId
		ID:        bson.NewObjectId(),
		Email:     nu.Email,
		UserName:  nu.UserName,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
		PhotoURL:  pu,
	}
	//call .SetPassword() to set the PassHash
	//field of the User to a hash of the NewUser.Password
	if err := u.SetPassword(nu.Password); err != nil {
		return nil, fmt.Errorf("Error setting password to hash: %v", err)
	}
	return u, nil
}

//FullName returns the user's full name, in the form:
//"<FirstName> <LastName>"
//If either first or last name is an empty string, no
//space is put betweeen the names
func (u *User) FullName() string {
	if len(u.FirstName) == 0 || len(u.LastName) == 0 {
		return u.FirstName + u.LastName
	}
	return u.FirstName + " " + u.LastName
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	//use the bcrypt package to generate a new hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return fmt.Errorf("Error generating bcrypt hash: %v", err)
	}
	u.PassHash = hash
	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	//use the bcrypt package to compare the supplied
	//password with the stored PassHash
	if err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password)); err != nil {
		return fmt.Errorf("Password doesn't match stored hash: %v", err)
	}
	return nil
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *Updates) error {
	//set the fields of `u` to the values of the related
	//field in the `updates` struct, enforcing the following rules:
	//the FirstName must be non-zero-length
	//the LastName must be non-zero-length
	if len(updates.FirstName) == 0 || len(updates.LastName) == 0 {
		return fmt.Errorf("Names must be non-zero length")
	}
	u.FirstName = updates.FirstName
	u.LastName = updates.LastName
	return nil
}

package indexes

import (
	"gopkg.in/mgo.v2/bson"
)

//Index outlines a key/value storage
type Index interface {
	AddUser(prefix string, userID bson.ObjectId)

	GetUsers(prefix string, limit int) []bson.ObjectId

	RemoveUser(prefix string, userID bson.ObjectId) bson.ObjectId
}

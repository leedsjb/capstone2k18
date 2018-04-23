package indexes

import (
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

//TestNewUserIdentifiers tests the NewUserIdentifiers function
func TestNewUserIdentifiers(t *testing.T) {
	set := NewUserIdentifiers()
	if set.col == nil {
		t.Error("Struct's map is not initialized")
	}
}

//TestAdd tests UserIdentifiers' Add function
func TestAdd(t *testing.T) {
	set := NewUserIdentifiers()

	userIDs := []bson.ObjectId{bson.NewObjectId()}

	cases := []struct {
		name      string
		ids       []bson.ObjectId
		expValues []bson.ObjectId
	}{
		{
			"unique id",
			userIDs,
			userIDs,
		},
		{
			"duplicate id",
			userIDs,
			userIDs,
		},
	}

	for _, c := range cases {
		for _, id := range c.ids {
			set.Add(id)
		}
		result := set.Sort()
		if !compareContents(c.expValues, result) {
			t.Errorf("Expected: %v, but got: %v", c.expValues, result)
		}
	}
}

//TestContains tests UserIdentifier's Contains function
func TestContains(t *testing.T) {
	userID1 := bson.NewObjectId()
	userID2 := bson.NewObjectId()

	set := NewUserIdentifiers()
	set.Add(userID1)

	cases := []struct {
		name     string
		id       bson.ObjectId
		expValue bool
	}{
		{
			"present ID",
			userID1,
			true,
		},
		{
			"non-present ID",
			userID2,
			false,
		},
	}

	for _, c := range cases {
		if contains := set.Contains(c.id); contains != c.expValue {
			t.Errorf("Expected: %v, but got: %v", c.expValue, contains)
		}
	}
}

//TestRemove tests UserIdentifier's Remove function
func TestRemove(t *testing.T) {
	userID1 := bson.NewObjectId()
	userID2 := bson.NewObjectId()

	set := NewUserIdentifiers()
	set.Add(userID1)

	cases := []struct {
		name      string
		id        bson.ObjectId
		expValues []bson.ObjectId
	}{
		{
			"non-present ID",
			userID2,
			[]bson.ObjectId{userID1},
		},
		{
			"present ID",
			userID1,
			[]bson.ObjectId{},
		},
	}

	for _, c := range cases {
		set.Remove(c.id)
		if result := set.Sort(); !compareContents(c.expValues, result) {
			t.Errorf("Expected: %v, but got :%v", c.expValues, result)
		}
	}
}

//TestSort tests UserIdentifier's Sort function
func TestSort(t *testing.T) {
	set := NewUserIdentifiers()
	userID1 := bson.ObjectId("5a013765cea32041111e97b0")
	userID2 := bson.ObjectId("5a013765cea32041111e97b1")
	set.Add(userID2)
	set.Add(userID1)

	cases := []struct {
		name      string
		expValues []bson.ObjectId
		expMatch  bool
	}{
		{
			"correct sort",
			[]bson.ObjectId{userID1, userID2},
			true,
		},
		{
			"incorrect sort",
			[]bson.ObjectId{userID2, userID1},
			false,
		},
	}

	for _, c := range cases {
		result := set.Sort()
		if !reflect.DeepEqual(c.expValues, result) && c.expMatch {
			t.Errorf("Error sorting. Expected: %v, but got: %v", c.expValues, result)
		}
	}
}

//TestSize tests UserIdentifiers's Size function
func TestSize(t *testing.T) {
	set := NewUserIdentifiers()
	userID1 := bson.NewObjectId()

	cases := []struct {
		name      string
		ids       []bson.ObjectId
		expLength int
	}{
		{
			"no add",
			[]bson.ObjectId{},
			0,
		},
		{
			"first unique add",
			[]bson.ObjectId{userID1},
			1,
		},
		{
			"duplicate add",
			[]bson.ObjectId{userID1},
			1,
		},
	}

	for _, c := range cases {
		for _, id := range c.ids {
			set.Add(id)
		}
		if length := set.Size(); c.expLength != length {
			t.Errorf("Lengths do not match. Expected: %v, but got: %v", c.expLength, length)
		}
	}
}

//TestRetainAll tests UserIdentifier's RetainAll function
func TestRetainAll(t *testing.T) {
	set1 := NewUserIdentifiers()
	set2 := NewUserIdentifiers()
	set3 := NewUserIdentifiers()

	userID1 := bson.NewObjectId()
	userID2 := bson.NewObjectId()
	userID3 := bson.NewObjectId()

	set1.Add(userID1)
	set1.Add(userID2)
	set2.Add(userID3)
	set3.Add(userID1)
	set3.Add(userID2)
	set3.Add(userID3)

	cases := []struct {
		name       string
		set1       *UserIdentifiers
		set2       *UserIdentifiers
		expOverlap []bson.ObjectId
	}{
		{
			"no overlap",
			set1,
			set2,
			[]bson.ObjectId{},
		},
		{
			"overlap",
			set1,
			set3,
			[]bson.ObjectId{userID1, userID2},
		},
	}

	for _, c := range cases {
		overlap := c.set1.RetainAll(c.set2)
		if !compareContents(c.expOverlap, overlap.Sort()) {
			t.Errorf("Expected: %v, but got: %v", c.expOverlap, overlap)
		}
	}
}

//compareContents is a helper function, written for these tests, that
//compares arrays while disregarding item order
func compareContents(users1, users2 []bson.ObjectId) bool {
	//https://stackoverflow.com/a/36000696
	if len(users1) != len(users2) {
		return false
	}
	diff := make(map[bson.ObjectId]int, len(users1))
	for _, user := range users1 {
		diff[user]++
	}
	for _, user := range users2 {
		if _, ok := diff[user]; !ok {
			return false
		}
		diff[user]--
		if diff[user] == 0 {
			delete(diff, user)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

//https://softwareengineering.stackexchange.com/a/177446

package indexes

import (
	"sort"

	"gopkg.in/mgo.v2/bson"
)

//UserIdentifiers is a collection of unique user IDs
type UserIdentifiers struct {
	col map[bson.ObjectId]bool
}

//NewUserIdentifiers makes a new UserIdentifiers
func NewUserIdentifiers() *UserIdentifiers {
	return &UserIdentifiers{make(map[bson.ObjectId]bool)}
}

//Add adds a new UserID to UserIdentifiers
func (ui *UserIdentifiers) Add(id bson.ObjectId) bool {
	contains := ui.col[id]
	ui.col[id] = true
	return !contains
}

//Contains checks if UserIdentifiers contains an ID
func (ui *UserIdentifiers) Contains(id bson.ObjectId) bool {
	return ui.col[id]
}

//Remove an ID from UserIdenitifiers
func (ui *UserIdentifiers) Remove(id bson.ObjectId) {
	delete(ui.col, id)
}

//Sort UserIdentifiers's IDs
func (ui *UserIdentifiers) Sort() []bson.ObjectId {
	tempSlice := make([]bson.ObjectId, len(ui.col))
	index := 0
	for id := range ui.col {
		tempSlice[index] = id
		index++
	}
	sort.Slice(tempSlice, func(i, j int) bool { return tempSlice[i] < tempSlice[j] })
	return tempSlice
}

//Size returns the number of unique IDs
func (ui *UserIdentifiers) Size() int {
	return len(ui.col)
}

//RetainAll returns a new UserIdentifiers with only the IDs contained in both sets
func (ui *UserIdentifiers) RetainAll(oui *UserIdentifiers) *UserIdentifiers {
	nui := NewUserIdentifiers()
	for id := range ui.col {
		if oui.Contains(id) {
			nui.Add(id)
		}
	}
	return nui
}

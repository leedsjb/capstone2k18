package indexes

import (
	"sort"
)

// EntityIdentifiers ...
type EntityIdentifiers struct {
	col map[int]bool
}

// NewEntityIdentifiers ...
func NewEntityIdentifiers() *EntityIdentifiers {
	return &EntityIdentifiers{make(map[int]bool)}
}

// Add ...
func (ei *EntityIdentifiers) Add(id int) bool {
	contains := ei.Contains(id)
	ei.col[id] = true
	return !contains
}

// Contains ...
func (ei *EntityIdentifiers) Contains(id int) bool {
	return ei.col[id]
}

// Remove ...
func (ei *EntityIdentifiers) Remove(id int) {
	delete(ei.col, id)
}

// Sort ...
func (ei *EntityIdentifiers) Sort() []int {
	tempSlice := make([]int, len(ei.col))
	index := 0
	for id := range ei.col {
		// We can also use append, but this might be more efficient
		tempSlice[index] = id
		index++
	}
	sort.Ints(tempSlice)
	return tempSlice
}

// Size ...
func (ei *EntityIdentifiers) Size() int {
	return len(ei.col)
}

// RetainAll ...
func (ei *EntityIdentifiers) RetainAll(oei *EntityIdentifiers) *EntityIdentifiers {
	nei := NewEntityIdentifiers()
	for id := range ei.col {
		if oei.Contains(id) {
			nei.Add(id)
		}
	}
	return nei
}

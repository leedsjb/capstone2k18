package indexes

//Index outlines a key/value storage
type Index interface {
	AddEntity(prefix string, entityID int)

	GetEntities(prefix string, limit int) []int

	RemoveEntity(prefix string, entityID int) int
}

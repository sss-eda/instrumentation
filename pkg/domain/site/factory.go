package site

// Factory TODO
type Factory interface {
	New() (ID, Aggregate, error)
}

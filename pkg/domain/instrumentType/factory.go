package instrumentType

// Factory TODO
type Factory interface {
	New() (ID, Aggregate, error)
}

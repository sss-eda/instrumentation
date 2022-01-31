package instrument

// Factory TODO
type Factory interface {
	New() (ID, Aggregate, error)
}

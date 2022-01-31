package instrument

// Storage TODO
type Storage interface {
	Load(ID) (Aggregate, error)
	Save(ID, Aggregate) error
}

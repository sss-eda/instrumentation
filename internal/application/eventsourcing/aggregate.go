package eventsourcing

// Aggregate TODO
type Aggregate interface{}

// AggregateID TODO
type AggregateID interface {
	Equals(AggregateID) bool
	String() string
}

// Entity TODO
type Entity interface{}

// Storage TODO
type Storage interface {
	Load(AggregateID) (Aggregate, error)
	Save(AggregateID, Aggregate) error
}

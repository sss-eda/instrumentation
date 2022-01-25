package commanding

// InstrumentID TODO
type InstrumentID interface {
	Equals(InstrumentID) bool
}

// Instrument TODO - Aggregate Root, Site and InstrumentType are entities.
type Instrument interface {
	Connect() (Connection, error)
}

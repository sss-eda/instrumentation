package instrumentation

// This is the service interface
type API interface {
	RegisterInstrument(string) error
	SendCommandToInstrument(string, Command) error
	GetAllInstrumentEvents(string) ([]*Event, error)
}

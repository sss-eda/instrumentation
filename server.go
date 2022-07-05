package instrumentation

type Server interface {
	RegisterInstrument(string) error
	SendCommandToInstrument(string, Command) error
	GetAllInstrumentEvents(string) ([]*Event, error)
}

type SendCommandToInstrument func(kind string, id string, command *Command) error

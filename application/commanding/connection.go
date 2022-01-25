package commanding

// CommandType TODO
type CommandType interface {
	String() string
}

// Command TODO
type Command interface {
	Type() CommandType
	Execute() error
}

// Connection TODO
type Connection interface {
	Transmit(Command) error
	Close() error
}

// InstrumentStorage TODO
type InstrumentStorage interface {
	Load(InstrumentID) (Instrument, error)
	Save(InstrumentID, Instrument) error
}

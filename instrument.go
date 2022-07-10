package instrumentation

type InstrumentID any

type Instrument interface {
	SendCommand(*Command) error
	GetAllEvents() ([]*Event, error)
	ApplyEvent(*Event) error
}

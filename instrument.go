package instrumentation

type Instrument struct {
	ID       string
	commands []*Command
	events   []*Event
}

func (instrument *Instrument) SendCommand(command *Command) error {
	instrument.commands = append(instrument.commands, command)

	return nil
}

func (instrument *Instrument) GetAllEvents() ([]*Event, error) {
	return instrument.events, nil
}

func (instrument *Instrument) ApplyEvent(event *Event) error {
	instrument.events = append(instrument.events, event)

	return nil
}

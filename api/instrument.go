package api

import "github.com/sss-eda/instrumentation"

type Instrument struct {
	ID       string
	commands []*instrumentation.Command
	events   []*instrumentation.Event
}

func (instrument *Instrument) SendCommand(command *instrumentation.Command) error {
	instrument.commands = append(instrument.commands, command)

	return nil
}

func (instrument *Instrument) GetAllEvents() ([]*instrumentation.Event, error) {
	return instrument.events, nil
}

func (instrument *Instrument) ApplyEvent(event *instrumentation.Event) error {
	instrument.events = append(instrument.events, event)

	return nil
}

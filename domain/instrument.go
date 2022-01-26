package domain

import "fmt"

// InstrumentID TODO
type InstrumentID interface {
	Equals(InstrumentID) bool
	String() string
}

// Instrument TODO
type Instrument struct {
	site           site
	instrumentType instrumentType
	commands       []Command
	iteration      uint8
	sequence       uint64
}

// NewInstrument TODO
func NewInstrument(
	site Site,
	instrumentType InstrumentType,
	iteration uint8,
) (*Instrument, error) {
	instrument := Instrument{
		site:           site,
		instrumentType: isntrumentType,
	}
}

// Add TODO
func (instrument *Instrument) Add(
	events chan<- Event,
) error {
	if instrument.sequence > 0 {
		return fmt.Errorf("Instrument already exists")
	}

	events <- InstrumentAddedEvent{
		SiteID:           instrument.site.ID(),
		InstrumentTypeID: instrument.instrumentType.ID(),
		Iteration:        instrument.iteration,
	}

	return nil
}

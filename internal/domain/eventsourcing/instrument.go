package eventsourcing

import (
	"github.com/sss-eda/instrumentation/internal/domain"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// Instrument TODO
type Instrument struct {
	id               domain.InstrumentID
	instrumentTypeID domain.InstrumentTypeID
	siteID           domain.SiteID
	name             domain.InstrumentName
	events           []Event
}

// NewInstrument TODO
func NewInstrument() *Instrument {
	return &Instrument{
		id:               NoInstrumentID{},
		instrumentTypeID: NoInstrumentTypeID{},
		siteID:           NoSiteID{},
		name:             domain.InstrumentName(""),
		events:           []Event{},
	}
}

// Activate TODO
func (instrument *Instrument) Activate() error {
	event := InstrumentActivatedEvent{
		InstrumentID: domain.Instrumentid,
	}

	instrument.changes = append(domain.Instrumentevents, event)

	return nil
}

// Deactivate TODO
func (instrument *Instrument) Deactivate() error {
	event := InstrumentDeactivatedEvent{
		InstrumentID: domain.Instrumentid,
	}

	domain.Instrumentevents = append(domain.Instrumentevents, event)

	return nil
}

// Relocate TODO
func (instrument *Instrument) Relocate(
	newSiteID site.ID,
) error {
	event := InstrumentRelocatedEvent{
		InstrumentID: domain.Instrumentid,
		NewSiteID:    newSiteID,
	}

	domain.Instrumentevents = append(domain.Instrumentevents, event)

	return nil
}

// Rename TODO
func (instrument *Instrument) Rename(
	newName domain.InstrumentName,
) error {
	event := InstrumentRenamedEvent{
		InstrumentID: domain.Instrumentid,
		NewName:      newName,
	}

	domain.Instrumentevents = append(domain.Instrumentevents, event)

	return nil
}

// ChangeInstrumentType TODO
func (instrument *Instrument) ChangeInstrumentType(
	newInstrumentTypeID domain.InstrumentTypeID,
) error {
	event := InstrumentTypeChangedEvent{
		InstrumentID:     instrument.ID,
		InstrumentTypeID: newInstrumentTypeID,
	}

	domain.Instrumentevents = append(domain.Instrumentevents, event)

	return nil
}

// NoInstrumentID TODO
type NoInstrumentID struct{}

// String TODO
func (NoInstrumentID) String() string {
	return "Currently not installed at any site."
}

// Equals TODO
func (NoInstrumentID) Equals(
	otherInstrumentID domain.InstrumentID,
) bool {
	var equals bool

	switch otherInstrumentID.(type) {
	case NoInstrumentID:
		equals = true
	default:
		equals = false
	}

	return equals
}

package domain

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// Instrument TODO
type Instrument struct {
	id               instrument.ID
	instrumentTypeID instrumentType.ID
	siteID           site.ID
	name             instrument.Name
	events           []Event
}

// NewInstrument TODO
func NewInstrument() instrument.Aggregate {
	return &Instrument{
		id:               NoInstrumentID{},
		instrumentTypeID: NoInstrumentTypeID{},
		siteID:           NoSiteID{},
		name:             instrument.Name(""),
		events:           []Event{},
	}
}

// Activate TODO
func (ar *Instrument) Activate() error {
	event := InstrumentActivatedEvent{
		InstrumentID: ar.id,
	}

	ar.events = append(ar.events, event)

	return nil
}

// Deactivate TODO
func (ar *Instrument) Deactivate() error {
	event := InstrumentDeactivatedEvent{
		InstrumentID: ar.id,
	}

	ar.events = append(ar.events, event)

	return nil
}

// Relocate TODO
func (ar *Instrument) Relocate(
	newSiteID site.ID,
) error {
	event := InstrumentRelocatedEvent{
		InstrumentID: ar.id,
		NewSiteID:    newSiteID,
	}

	ar.events = append(ar.events, event)

	return nil
}

// Rename TODO
func (ar *Instrument) Rename(
	newName instrument.Name,
) error {
	event := InstrumentRenamedEvent{
		InstrumentID: ar.id,
		NewName:      newName,
	}

	ar.events = append(ar.events, event)

	return nil
}

// ChangeInstrumentType TODO
func (ar *Instrument) ChangeInstrumentType(
	instrumentTypeID instrumentType.ID,
) error {
	event := InstrumentTypeChangedEvent{
		InstrumentID:        ar.id,
		NewInstrumentTypeID: instrumentTypeID,
	}

	ar.events = append(ar.events, event)

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
	otherInstrumentID instrument.ID,
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

package domain

import "fmt"

// Site - This will be the aggregate root
type Site struct {
	id           SiteID
	instruments  map[InstrumentTypeID][]Instrument
	name         string
	abbreviation string
	events       chan Event
}

// CommissionInstrument TODO - This should just be add and remove. The application level will have to decide on commission, move or decommission.
func (site Site) CommissionInstrument(
	instrumentID InstrumentID,
	instrumentTypeID InstrumentTypeID,
) error {
	var iteration uint8 = 0
	instruments := site.instruments[instrumentTypeID]
	for _, instrument := range instruments {
		if instrument.iteration > iteration {
			iteration = instrument.iteration
		}
	}
	iteration++

	event := InstrumentCommissionedEvent{
		InstrumentID:     instrumentID,
		InstrumentTypeID: instrumentTypeID,
		Iteration:        iteration,
	}

	site.events <- event

	return nil
}

// DecommissionInstrument TODO
func (site Site) DecommissionInstrument(
	instrumentTypeID InstrumentTypeID,
	iteration uint8,
) error {
	instruments := site.instruments[instrumentTypeID]

	var instrumentID InstrumentID = NoInstrumentID{}

	for _, instrument := range instruments {
		if instrument.iteration == iteration {
			instrumentID = instrument.id
			break
		}
	}

	if instrumentID.Equals(NoInstrumentID{}) {
		return fmt.Errorf(
			"no instrument with type id: %v of iteration: %d at site: %s",
			instrumentTypeID,
			iteration,
			site.name,
		)
	}

	event := InstrumentDecommissionedEvent{
		InstrumentID: instrumentID,
	}

	site.events <- event

	return nil
}

// SiteID TODO
type SiteID interface {
	Equals(SiteID) bool
	String() string
}

// NoSiteID TODO
type NoSiteID struct{}

// String TODO
func (NoSiteID) String() string {
	return "Currently not installed at any site."
}

// Equals TODO
func (NoSiteID) Equals(
	otherSiteID SiteID,
) bool {
	var equals bool

	switch otherSiteID.(type) {
	case NoSiteID:
		equals = true
	default:
		equals = false
	}

	return equals
}

// ID TODO
func (site Site) ID() SiteID {
	return site.id
}

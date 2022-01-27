package domain

// Instrument TODO
type Instrument struct {
	id        InstrumentID
	iteration uint8
}

// InstrumentID TODO
type InstrumentID interface {
	Equals(InstrumentID) bool
	String() string
}

// NoInstrumentID TODO
type NoInstrumentID struct{}

// String TODO
func (NoInstrumentID) String() string {
	return "Currently not installed at any site."
}

// Equals TODO
func (NoInstrumentID) Equals(
	otherInstrumentID InstrumentID,
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

// Relocate TODO
// func (instrument *Instrument) Relocate(
// 	events chan<- Event,
// 	newSite Site,
// ) error {
// 	if newSite.ID().Equals(instrument.SiteID) {
// 		return fmt.Errorf("Instrument already at this site")
// 	}

// 	event := InstrumentRelocatedEvent{
// 		SiteID: newSiteID,
// 	}

// 	events <- event

// 	return nil
// }

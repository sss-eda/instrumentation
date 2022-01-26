package domain

// Event TODO
type Event interface{}

// InstrumentAddedEvent TODO
type InstrumentAddedEvent struct {
	SiteID           SiteID
	InstrumentTypeID InstrumentTypeID
	Iteration        uint8
}

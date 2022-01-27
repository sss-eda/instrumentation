package domain

// Event TODO
type Event interface{}

// InstrumentCommissionedEvent TODO
type InstrumentCommissionedEvent struct {
	InstrumentID     InstrumentID
	InstrumentTypeID InstrumentTypeID
	Iteration        uint8
}

// InstrumentDecommissionedEvent TODO
type InstrumentDecommissionedEvent struct {
	InstrumentID InstrumentID
}

// InstrumentRelocatedEvent TODO
type InstrumentRelocatedEvent struct {
	SiteID SiteID
}

// InstrumentTypeRenamedEvent TODO
type InstrumentTypeRenamedEvent struct {
	NewName string
}

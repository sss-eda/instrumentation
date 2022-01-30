package domain

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// Event TODO
type Event interface{}

// InstrumentCommissionedEvent TODO
type InstrumentCommissionedEvent struct {
	InstrumentID     instrument.ID
	InstrumentTypeID instrumentType.ID
	Iteration        uint8
}

type SiteAbbreviationChangedEvent struct {
	NewAbbreviation site.Abbreviation
}

// InstrumentRelocatedEvent TODO
type InstrumentRelocatedEvent struct {
	InstrumentID instrument.ID
	NewSiteID    site.ID
}

// InstrumentTypeRenamedEvent TODO
type InstrumentTypeRenamedEvent struct {
	InstrumentTypeID instrumentType.ID
	NewName          instrumentType.Name
}

// SiteRenamedEvent TODO
type SiteRenamedEvent struct {
	SiteID  site.ID
	NewName site.Name
}

// InstrumentActivatedEvent TODO
type InstrumentActivatedEvent struct {
	InstrumentID instrument.ID
}

// InstrumentDeactivatedEvent TODO
type InstrumentDeactivatedEvent struct {
	InstrumentID instrument.ID
}

// InstrumentRenamedEvent TODO
type InstrumentRenamedEvent struct {
	InstrumentID instrument.ID
	NewName      instrument.Name
}

// InstrumentTypeChangedEvent TODO
type InstrumentTypeChangedEvent struct {
	InstrumentID        instrument.ID
	NewInstrumentTypeID instrumentType.ID
}

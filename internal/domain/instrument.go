package domain

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// Instrument TODO
type Instrument interface {
	Rename(Name) error
	Relocate(site.ID) error
	ChangeInstrumentType(instrumentType.ID) error
	Deactivate() error
	Activate() error
}

// InstrumentID TODO
type InstrumentID interface {
	Equals(InstrumentID) bool
	String() string
}

// InstrumentName TODO
type InstrumentName string

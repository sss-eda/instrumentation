package domain

import (
	"fmt"

	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
)

// InstrumentType TODO
type InstrumentType struct {
	id           instrumentType.ID
	name         instrumentType.Name
	abbreviation instrumentType.Abbreviation
}

// Rename TODO
func (ar InstrumentType) Rename(
	newName instrumentType.Name,
) error {
	if len(newName) < 1 {
		return fmt.Errorf("name may not be an empty string")
	}

	ar.name = newName

	return nil
}

// ChangeAbbreviation TODO
func (ar *InstrumentType) ChangeAbbreviation(
	newAbbreviation instrumentType.Abbreviation,
) error {
	ar.abbreviation = newAbbreviation

	return nil
}

// NoInstrumentTypeID TODO
type NoInstrumentTypeID struct{}

// String TODO
func (NoInstrumentTypeID) String() string {
	return "Currently not installed at any site."
}

// Equals TODO
func (NoInstrumentTypeID) Equals(
	otherInstrumentTypeID instrumentType.ID,
) bool {
	var equals bool

	switch otherInstrumentTypeID.(type) {
	case NoInstrumentTypeID:
		equals = true
	default:
		equals = false
	}

	return equals
}

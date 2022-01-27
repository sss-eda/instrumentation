package domain

import "fmt"

// InstrumentType TODO
type InstrumentType struct {
	id           InstrumentTypeID
	name         string
	abbreviation string
	events       chan Event
}

// InstrumentTypeID TODO
type InstrumentTypeID interface {
	Equals(InstrumentTypeID) bool
	String() string
}

// NoInstrumentTypeID TODO
type NoInstrumentTypeID struct{}

// String TODO
func (NoInstrumentTypeID) String() string {
	return "Currently not installed at any site."
}

// Equals TODO
func (NoInstrumentTypeID) Equals(
	otherInstrumentTypeID InstrumentTypeID,
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

// ID TODO
func (instrumentType InstrumentType) ID() InstrumentTypeID {
	return instrumentType.id
}

// Rename TODO
func (instrumentType InstrumentType) Rename(
	newName string,
) error {
	if len(newName) < 1 {
		return fmt.Errorf("name may not be an empty string")
	}

	event := InstrumentTypeRenamedEvent{
		NewName: newName,
	}

	instrumentType.events <- event

	return nil
}

package adding

// InstrumentTypeID TODO
type InstrumentTypeID interface {
	Equals(InstrumentTypeID) bool
}

// InstrumentType TODO
type InstrumentType struct {
	Name         string
	Abbreviation string
}

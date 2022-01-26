package domain

// InstrumentTypeID TODO
type InstrumentTypeID interface {
	Equals(InstrumentTypeID) bool
	String() string
}

// InstrumentType TODO
type instrumentType struct {
	id           InstrumentTypeID
	name         string
	abbreviation string
}

// ID TODO
func (entity instrumentType) ID() InstrumentTypeID {
	return entity.id
}

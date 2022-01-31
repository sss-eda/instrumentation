package domain

// InstrumentType TODO
type InstrumentType interface {
	Rename(InstrumentTypeName) error
	ChangeAbbreviation(InstrumentTypeAbbreviation) error
}

// InstrumentTypeID - Maybe move to different layer?
type InstrumentTypeID interface {
	Equals(InstrumentTypeID) bool
	String() string
}

// InstrumentTypeName TODO
type InstrumentTypeName string

// InstrumentTypeAbbreviation TODO
type InstrumentTypeAbbreviation string

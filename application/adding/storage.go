package adding

// InstrumentStorage TODO
type InstrumentStorage interface {
	Add(InstrumentID, Instrument) error
}

// SiteStorage TODO
type SiteStorage interface {
	Add(SiteID, Site) error
}

// InstrumentTypeStorage TODO
type InstrumentTypeStorage interface {
	Add(InstrumentTypeID, InstrumentType) error
}

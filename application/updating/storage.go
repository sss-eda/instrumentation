package adding

// InstrumentStorage TODO
type InstrumentStorage interface {
	Update(InstrumentID, Instrument) error
}

// SiteStorage TODO
type SiteStorage interface {
	Update(SiteID, Site) error
}

// InstrumentTypeStorage TODO
type InstrumentTypeStorage interface {
	Update(InstrumentTypeID, InstrumentType) error
}

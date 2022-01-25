package application

import "github.com/sss-eda/administration/internal/application/adding"

// Adder TODO
type Adder interface {
	AddInstrument(adding.InstrumentID, adding.Instrument) error
	AddInstrumentType(adding.InstrumentTypeID, adding.InstrumentType) error
	AddSite(adding.SiteID, adding.Site) error
}

// Updater TODO
type Updater interface {
	UpdateInstrument(updating.InstrumentID, updating.Instrument) error
	UpdateInstrumentType(updating.InstrumentTypeID, updating.InstrumentType) error
	UpdateSite(updating.SiteID, updating.Site) error
}

// Finder TODO
type Finder interface {
	FindInstrumentByAbbreviation(string) finding.Instrument
	FindInstrumentTypeByAbbreviation(string) finding.InstrumentType
	FindSiteByAbbreviation(string) finding.InstrumentType
}

// Lister TODO
type Lister interface {
	ListInstruments() ([]listing.Instrument, error)
	ListInstrumentTypes() ([]listing.InstrymentType, error)
	ListSites() ([]listing.Site, error)
}

package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// Service TODO
type Service interface {
	AddInstrument(instrument.Name, instrumentType.ID, site.ID)
	RelocateInstrument(instrument.ID, site.ID)
	ChangeInstrumentType(instrument.ID, instrumentType.ID)
	ActivateInstrument(instrument.ID)
	DeactivateInstrument(instrument.ID)
	RenameInstrument(instrument.ID, instrument.Name)
	AddSite(site.Name, site.Abbreviation)
	RemoveSite(site.ID)
	RenameSite(site.Name)
	ChangeSiteAbbreviation(site.ID, site.Abbreviation)
	AddInstrumentType(instrumentType.Name, instrumentType.Abbreviation)
	RenameInstrumentType(instrumentType.ID, instrumentType.Name)
	ChangeInstrumentTypeAbbreviation(instrumentType.Abbreviation)
}

package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// InstrumentRepository TODO
type InstrumentRepository interface {
	New() (instrument.ID, instrument.Aggregate)
	Load(instrument.ID) (instrument.Aggregate, error)
	Save(instrument.ID, instrument.Aggregate) error
}

// SiteRepository TODO
type SiteRepository interface {
	New() (site.ID, site.Aggregate)
	Load(site.ID) (site.Aggregate, error)
	Save(site.ID, site.Aggregate) error
}

// InstrumentTypeRepository TODO
type InstrumentTypeRepository interface {
	New() (instrumentType.ID, instrumentType.Aggregate)
	Load(instrumentType.ID) (instrumentType.Aggregate, error)
	Save(instrumentType.ID, instrumentType.Aggregate) error
}

package instrument

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// Aggregate TODO
type Aggregate interface {
	Add(Name, site.ID, instrumentType.ID) error
	Rename(Name) error
	Relocate(site.ID) error
	ChangeInstrumentType(instrumentType.ID) error
	Deactivate() error
	Activate() error
	Configure(Configuration) error
}

package jetstream

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// InstrumentRepository TODO
type InstrumentRepository struct{}

// New TODO
func (repository *InstrumentRepository) New() (
	instrument.ID,
	instrument.Aggregate,
) {
	return nil, nil
}

// Load TODO
func (repository *InstrumentRepository) Load(
	id site.ID,
) (site.Aggregate, error) {
	return nil, nil
}

// Load TODO
func (repository *InstrumentRepository) Save(
	id site.ID,
	aggregate site.Aggregate,
) error {
	return nil
}

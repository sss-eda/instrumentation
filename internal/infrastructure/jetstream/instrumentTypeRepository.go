package jetstream

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// InstrumentTypeRepository TODO
type InstrumentTypeRepository struct{}

// New TODO
func (repository *InstrumentTypeRepository) New() (
	instrumentType.ID,
	instrumentType.Aggregate,
) {
	return nil, nil
}

// Load TODO
func (repository *InstrumentTypeRepository) Load(
	id instrumentType.ID,
) (instrumentType.Aggregate, error) {
	return nil, nil
}

// Load TODO
func (repository *InstrumentTypeRepository) Save(
	id site.ID,
	aggregate site.Aggregate,
) error {
	return nil
}

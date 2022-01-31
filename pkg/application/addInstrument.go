package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// AddInstrumentInput TODO
type AddInstrumentInput struct {
	InstrumentName   instrument.Name
	InstrumentTypeID instrumentType.ID
	SiteID           site.ID
}

// AddInstrumentMutation TODO
func AddInstrumentMutation(
	factory instrument.Factory,
	storage instrument.Storage,
) (
	func(AddInstrumentInput) error,
	error,
) {
	return func(input AddInstrumentInput) error {
		id, aggregate, err := factory.New()
		if err != nil {
			return err
		}

		err = aggregate.Add(
			input.InstrumentName,
			input.SiteID,
			input.InstrumentTypeID,
		)
		if err != nil {
			return err
		}

		err = storage.Save(id, aggregate)
		if err != nil {
			return err
		}

		return nil
	}, nil
}

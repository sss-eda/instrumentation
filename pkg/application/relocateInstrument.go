package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// RelocateInstrumentInput TODO
type RelocateInstrumentInput struct {
	InstrumentID instrument.ID
	NewSiteID    site.ID
}

// RelocateInstrumentMutation TODO
func RelocateInstrumentMutation(
	storage instrument.Storage,
) (
	func(RelocateInstrumentInput) error,
	error,
) {
	return func(input RelocateInstrumentInput) error {
		aggregate, err := storage.Load(input.InstrumentID)
		if err != nil {
			return err
		}

		err = aggregate.Relocate(
			input.NewSiteID,
		)
		if err != nil {
			return err
		}

		err = storage.Save(input.InstrumentID, aggregate)
		if err != nil {
			return err
		}

		return nil
	}, nil
}

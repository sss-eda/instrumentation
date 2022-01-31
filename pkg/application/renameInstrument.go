package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
)

// RenameInstrumentInput TODO
type RenameInstrumentInput struct {
	InstrumentID      instrument.ID
	NewInstrumentName instrument.Name
}

// RenameInstrumentMutation TODO
func RenameInstrumentMutation(
	storage instrument.Storage,
) (
	func(RenameInstrumentInput) error,
	error,
) {
	return func(input RenameInstrumentInput) error {
		aggregate, err := storage.Load(input.InstrumentID)
		if err != nil {
			return err
		}

		err = aggregate.Rename(
			input.NewInstrumentName,
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

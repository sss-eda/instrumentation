package domain

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
)

// ChangeInstrumentTypeInput TODO
type ChangeInstrumentTypeInput struct {
	InstrumentID        instrument.ID
	NewInstrumentTypeID instrumentType.ID
}

// ChangeInstrumentTypeMutation TODO
func ChangeInstrumentTypeMutation(
	storage instrument.Storage,
) (
	func(ChangeInstrumentTypeInput) error,
	error,
) {
	return func(input ChangeInstrumentTypeInput) error {
		aggregate, err := storage.Load(input.InstrumentID)
		if err != nil {
			return err
		}

		err = aggregate.ChangeInstrumentType(
			input.NewInstrumentTypeID,
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

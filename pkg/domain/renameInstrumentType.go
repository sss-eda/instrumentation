package domain

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
)

// RenameInstrumentTypeInput TODO
type RenameInstrumentTypeInput struct {
	InstrumentTypeID      instrumentType.ID
	NewInstrumentTypeName instrumentType.Name
}

// RenameInstrumentTypeMutation TODO
func RenameInstrumentTypeMutation(
	storage instrumentType.Storage,
) (
	func(RenameInstrumentTypeInput) error,
	error,
) {
	return func(input RenameInstrumentTypeInput) error {
		aggregate, err := storage.Load(input.InstrumentTypeID)
		if err != nil {
			return err
		}

		err = aggregate.Rename(
			input.NewInstrumentTypeName,
		)
		if err != nil {
			return err
		}

		err = storage.Save(input.InstrumentTypeID, aggregate)
		if err != nil {
			return err
		}

		return nil
	}, nil
}

package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
)

// ChangeInstrumentTypeAbbreviationInput TODO
type ChangeInstrumentTypeAbbreviationInput struct {
	InstrumentTypeID              instrumentType.ID
	NewInstrumentTypeAbbreviation instrumentType.Abbreviation
}

// ChangeInstrumentTypeAbbreviationMutation TODO
func ChangeInstrumentTypeAbbreviationMutation(
	storage instrumentType.Storage,
) (
	func(ChangeInstrumentTypeAbbreviationInput) error,
	error,
) {
	return func(input ChangeInstrumentTypeAbbreviationInput) error {
		aggregate, err := storage.Load(input.InstrumentTypeID)
		if err != nil {
			return err
		}

		err = aggregate.ChangeAbbreviation(
			input.NewInstrumentTypeAbbreviation,
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

package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
)

// AddInstrumentTypeInput TODO
type AddInstrumentTypeInput struct {
	InstrumentTypeName         instrumentType.Name
	InstrumentTypeAbbreviation instrumentType.Abbreviation
}

// AddInstrumentTypeMutation TODO
func AddInstrumentTypeMutation(
	factory instrumentType.Factory,
	storage instrumentType.Storage,
) (
	func(AddInstrumentTypeInput) error,
	error,
) {
	return func(input AddInstrumentTypeInput) error {
		id, aggregate, err := factory.New()
		if err != nil {
			return err
		}

		err = aggregate.Add(
			input.InstrumentTypeName,
			input.InstrumentTypeAbbreviation,
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

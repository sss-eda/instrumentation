package domain

import "github.com/sss-eda/instrumentation/pkg/domain/instrument"

// DectivateInstrumentInput TODO
type DectivateInstrumentInput struct {
	InstrumentID instrument.ID
}

// DeactivateInstrumentMutation TODO
func DeactivateInstrumentMutation(
	storage instrument.Storage,
) (
	func(DectivateInstrumentInput) error,
	error,
) {
	return func(input DectivateInstrumentInput) error {
		aggregate, err := storage.Load(input.InstrumentID)
		if err != nil {
			return err
		}

		err = aggregate.Activate()
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

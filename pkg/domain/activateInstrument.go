package domain

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
)

// ActivateInstrumentInput TODO
type ActivateInstrumentInput struct {
	InstrumentID instrument.ID
}

// InstrumentActivator TODO
type InstrumentActivator interface {
	ActivateInstrument(ActivateInstrumentInput) error
}

// ActivateInstrumentMutation TODO
func ActivateInstrumentMutation(
	storage instrument.Storage,
) (
	instrumentation.InstrumentActivator,
	error,
) {
	return func(input instrumentation.ActivateInstrumentInput) error {
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

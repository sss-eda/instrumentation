package instrumentation

import (
	"github.com/sss-eda/instrumentation/internal/domain/eventsourcing"

	"github.com/sss-eda/instrumentation/pkg/domain"
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
)

// Instrument TODO
type Instrument struct {
	Root instrument.Entity
}

// NewInstrument TODO
func NewInstrument(
	options ...InstrumentOption,
) (domain.Instrument, error) {
	instrument := Instrument{}

	var err error
	for _, option := range options {
		err = option(&instrument)
		if err != nil {
			return nil, err
		}
	}

	if instrument.Root == nil {
		instrument.Root = plain.NewInstrument()
	}

	return &instrument, nil
}

// InstrumentOption TODO
type InstrumentOption func(*Instrument) error

// WithEventSourcing TODO
func WithEventSourcing() func(*Instrument) error {
	return func(instrument *Instrument) error {
		instrument.Root = eventsourcing.NewInstrument()
	}
}

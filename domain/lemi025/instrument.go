package lemi025

// Instrument TODO
type Instrument struct {
	Config Config
}

// Mutate TODO
func (instrument *Instrument) Mutate(
	mutation MutationType,
) func() error {
	switch mutation {
	case ReadConfig:
		return instrument.Config.Read
	case ReadTime:
		return instrument.Time.Read
	case SetTime:
		return instrument.Time.Set
	}

	return nil
}

// Subscribe TODO
func (instrument *Instrument) Subscribe(
	event EventType,
) {
	switch event {
	case ConfigRead:
	case TimeRead:
	case TimeSet:
	}
}

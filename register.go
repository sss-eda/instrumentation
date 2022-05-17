package instrumentation

// Register TODO
func Register(repository Repository) func(string) error {
	return func(id string) error {
		instrument, err := repository.GetInstrumentByID(id)
		if err != nil {
			return err
		}

		return instrument.Register()
	}
}

package commanding

// Service TODO
type Service struct {
	instruments InstrumentStorage
}

// SendCommand TODO
func (service *Service) SendCommand(
	instrumentID InstrumentID,
	command Command,
) error {
	instrument, err := service.instruments.Load(instrumentID)
	if err != nil {
		return err
	}

	connection, err := instrument.Connect()
	if err != nil {
		return err
	}
	defer connection.Close()

	err = connection.Transmit(command)
	if err != nil {
		return err
	}

	return nil
}

package server

import "github.com/sss-eda/instrumentation"

func SendCommandToInstrument(
	getInstrumentByID instrumentation.GetInstrumentByID,
) instrumentation.SendCommandToInstrumentUseCase {
	return func(kind string, id string, command *instrumentation.Command) error {
		instrument, err := getInstrumentByID(id)
		if err != nil {
			return err
		}

		return instrument.SendCommand(command)
	}
}

package api

import "github.com/sss-eda/instrumentation"

type InstrumentRepository interface {
	GetInstrumentByID(string) (*Instrument, error)
	CreateInstrument(string) error
}

type Service struct {
	Instruments InstrumentRepository
}

func New(instruments InstrumentRepository) (*Service, error) {
	return &Service{
		Instruments: instruments,
	}, nil
}

func (service *Service) RegisterInstrument(id string) error {
	return service.Instruments.CreateInstrument(id)
}

func (service *Service) SendCommandToInstrument(id string, command *instrumentation.Command) error {
	instrument, err := service.Instruments.GetInstrumentByID(id)
	if err != nil {
		return err
	}

	err = instrument.SendCommand(command)

	return err
}

func (service *Service) GetAllInstrumentEvents(id string) ([]*instrumentation.Event, error) {
	instrument, err := service.Instruments.GetInstrumentByID(id)
	if err != nil {
		return nil, err
	}

	events, err := instrument.GetAllEvents()

	return events, err
}

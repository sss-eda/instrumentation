package server

import "github.com/sss-eda/instrumentation"

type Repository interface {
	GetInstrumentByID(string) (*instrumentation.Instrument, error)
	CreateInstrument(string) error
}

type Service struct {
	Repository
}

func (service *Service) RegisterInstrument(id string) error {
	return service.Repository.CreateInstrument(id)
}

func (service *Service) SendCommandToInstrument(id string, command *instrumentation.Command) error {
	instrument, err := service.Repository.GetInstrumentByID(id)
	if err != nil {
		return err
	}

	err = instrument.SendCommand(command)

	return err
}

func (service *Service) GetAllInstrumentEvents(id string) ([]*instrumentation.Event, error) {
	instrument, err := service.Repository.GetInstrumentByID(id)
	if err != nil {
		return nil, err
	}

	events, err := instrument.GetAllEvents()

	return events, err
}

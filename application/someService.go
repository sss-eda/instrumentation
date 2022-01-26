package application

import "github.com/sss-eda/instrumentation/domain"

// SomeApplicationService TODO
type SomeApplicationService interface{}

// Repository TODO
type SiteStorage interface {
}

type someApplicationService struct {
	// ds   domain.SomeDomainService
	repo   Repository
	events chan domain.Event
}

// AddInstrument TODO
func (service *someApplicationService) AddInstrument(
	siteID domain.SiteID,
	typeID domain.InstrumentTypeID,
	iteration uint8,
) error {
	site, err := service.repo.Load(siteID)
	if err != nil {
		return err
	}

	instrumentType, err := service.repo.Load(instrumentTypeID)
	if err != nil {
		return err
	}

	instrument, err := domain.NewInstrument(
		site,
		instrumentType,
		iteration,
	)
	if err != nil {
		return err
	}

	err = instrument.Add(service.events)
	if err != nil {
		return err
	}

	return err
}

package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// Service TODO
type Service struct {
	instruments     InstrumentRepository
	sites           SiteRepository
	instrumentTypes InstrumentTypeRepository
}

// NewService TODO
func NewService(
	instrumentRepository InstrumentRepository,
	siteRepository SiteRepository,
	instrumentTypeRepository InstrumentTypeRepository,
) (*Service, error) {
	return &Service{
		instruments:     instrumentRepository,
		sites:           siteRepository,
		instrumentTypes: instrumentTypeRepository,
	}, nil
}

// AddInstrument TODO
func (service *Service) AddInstrument(
	instrumentName instrument.Name,
	instrumentTypeID instrumentType.ID,
	siteID site.ID,
) error {
	myInstrumentID, myInstrument := service.instruments.New()

	err := service.instruments.Save(myInstrumentID, myInstrument)
	if err != nil {
		return err
	}

	return nil
}

// RelocateInstrument TODO
func (service *Service) RelocateInstrument(
	instrumentID instrument.ID,
	newSiteID site.ID,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrument, err := service.instruments.Load(instrumentID)
	if err != nil {
		return err
	}

	err = myInstrument.Relocate(newSiteID)
	if err != nil {
		return err
	}

	err = service.instruments.Save(instrumentID, myInstrument)
	if err != nil {
		return err
	}

	return nil
}

// ChangeInstrumentType TODO
func (service *Service) ChangeInstrumentType(
	instrumentID instrument.ID,
	newInstrumentTypeID instrumentType.ID,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrument, err := service.instruments.Load(instrumentID)
	if err != nil {
		return err
	}

	err = myInstrument.ChangeInstrumentType(newInstrumentTypeID)
	if err != nil {
		return err
	}

	err = service.instruments.Save(instrumentID, myInstrument)
	if err != nil {
		return err
	}

	return nil
}

// ActivateInstrument TODO
func (service *Service) ActivateInstrument(
	instrumentID instrument.ID,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrument, err := service.instruments.Load(instrumentID)
	if err != nil {
		return err
	}

	err = myInstrument.Activate()
	if err != nil {
		return err
	}

	err = service.instruments.Save(instrumentID, myInstrument)
	if err != nil {
		return err
	}

	return nil
}

// DeactivateInstrument TODO
func (service *Service) DeactivateInstrument(
	instrumentID instrument.ID,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrument, err := service.instruments.Load(instrumentID)
	if err != nil {
		return err
	}

	err = myInstrument.Deactivate()
	if err != nil {
		return err
	}

	err = service.instruments.Save(instrumentID, myInstrument)
	if err != nil {
		return err
	}

	return nil
}

// DeactivateInstrument TODO
func (service *Service) RenameInstrument(
	instrumentID instrument.ID,
	newName instrument.Name,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrument, err := service.instruments.Load(instrumentID)
	if err != nil {
		return err
	}

	err = myInstrument.Rename(newName)
	if err != nil {
		return err
	}

	err = service.instruments.Save(instrumentID, myInstrument)
	if err != nil {
		return err
	}

	return nil
}

// AddSite TODO
func (service *Service) AddSite(
	siteName site.Name,
	siteAbbreviation site.Abbreviation,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrumentID, myInstrument := service.instruments.New()

	err := service.instruments.Save(myInstrumentID, myInstrument)
	if err != nil {
		return err
	}

	return nil
}

// RemoveSite TODO
func (service *Service) RemoveSite(
	siteID site.ID,
) error {
	// Eewww! A variable with "my" in front of it.
	mySite, err := service.sites.Load(siteID)
	if err != nil {
		return err
	}

	err = mySite.Remove()
	if err != nil {
		return err
	}

	err = service.sites.Save(siteID, mySite)
	if err != nil {
		return err
	}

	return nil
}

// RenameSite TODO
func (service *Service) RenameSite(
	siteID site.ID,
	siteName site.Name,
) error {
	// Eewww! A variable with "my" in front of it.
	mySite, err := service.sites.Load(siteID)
	if err != nil {
		return err
	}

	err = mySite.Rename(siteName)
	if err != nil {
		return err
	}

	err = service.sites.Save(siteID, mySite)
	if err != nil {
		return err
	}

	return nil
}

// ChangeSiteAbbreviation TODO
func (service *Service) ChangeSiteAbbreviation(
	siteID site.ID,
	newAbbreviation site.Abbreviation,
) error {
	// Eewww! A variable with "my" in front of it.
	mySite, err := service.sites.Load(siteID)
	if err != nil {
		return err
	}

	err = mySite.ChangeAbbreviation(newAbbreviation)
	if err != nil {
		return err
	}

	err = service.sites.Save(siteID, mySite)
	if err != nil {
		return err
	}

	return nil
}

// AddInstrumentType TODO
func (service *Service) AddInstrumentType(
	name instrumentType.Name,
	abbreviation instrumentType.Abbreviation,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrumentTypeID, myInstrumentType := service.instrumentTypes.New()

	err := service.instrumentTypes.Save(myInstrumentTypeID, myInstrumentType)
	if err != nil {
		return err
	}

	return nil
}

// RenameInstrumentType TODO
func (service *Service) RenameInstrumentType(
	instrumentTypeID instrumentType.ID,
	newName instrumentType.Name,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrumentType, err := service.instrumentTypes.Load(instrumentTypeID)
	if err != nil {
		return err
	}

	err = myInstrumentType.Rename(newName)
	if err != nil {
		return err
	}

	err = service.instrumentTypes.Save(instrumentTypeID, myInstrumentType)
	if err != nil {
		return err
	}

	return nil
}

// ChangeInstrumentTypeAbbreviation TODO
func (service *Service) ChangeInstrumentTypeAbbreviation(
	instrumentTypeID instrumentType.ID,
	newAbbreviation instrumentType.Abbreviation,
) error {
	// Eewww! A variable with "my" in front of it.
	myInstrumentType, err := service.instrumentTypes.Load(instrumentTypeID)
	if err != nil {
		return err
	}

	err = myInstrumentType.ChangeAbbreviation(newAbbreviation)
	if err != nil {
		return err
	}

	err = service.instrumentTypes.Save(instrumentTypeID, myInstrumentType)
	if err != nil {
		return err
	}

	return nil
}

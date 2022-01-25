package adding

// Service TODO
type Service struct {
	instruments     InstrumentStorage
	instrumentTypes InstrumentTypeStorage
	sites           SiteStorage
}

// UpdateInstrument TODO
func (service *Service) UpdateInstrument(
	id InstrumentID,
	instrument Instrument,
) error {
	err := service.instruments.Update(id, instrument)
	if err != nil {
		return err
	}

	return nil
}

// UpdateInstrumentType TODO
func (service *Service) UpdateInstrumentType(
	id InstrumentTypeID,
	instrumentType InstrumentType,
) error {
	err := service.instrumentTypes.Update(id, instrumentType)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSite TODO
func (service *Service) UpdateSite(
	id SiteID,
	site Site,
) error {
	err := service.sites.Update(id, site)
	if err != nil {
		return err
	}

	return nil
}

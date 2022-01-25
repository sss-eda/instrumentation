package adding

// Service TODO
type Service struct {
	instruments     InstrumentStorage
	instrumentTypes InstrumentTypeStorage
	sites           SiteStorage
}

// AddInstrument TODO
func (service *Service) AddInstrument(
	id InstrumentID,
	instrument Instrument,
) error {
	err := service.instruments.Add(id, instrument)
	if err != nil {
		return err
	}

	return nil
}

// AddInstrumentType TODO
func (service *Service) AddInstrumentType(
	id InstrumentTypeID,
	instrumentType InstrumentType,
) error {
	err := service.instrumentTypes.Add(id, instrumentType)
	if err != nil {
		return err
	}

	return nil
}

// AddSite TODO
func (service *Service) AddSite(
	id SiteID,
	site Site,
) error {
	err := service.sites.Add(id, site)
	if err != nil {
		return err
	}

	return nil
}

package application

// // SomeApplicationService TODO
// type SomeApplicationService interface{}

// // SiteStorage TODO
// type SiteStorage interface {
// 	Load(domain.SiteID) (domain.Site, error)
// 	Save(domain.SiteID, domain.Site) error
// }

// // InstrumentTypeStorage TODO
// type InstrumentTypeStorage interface {
// 	Load(domain.InstrumentTypeID) (domain.InstrumentType, error)
// 	Save(domain.InstrumentTypeID, domain.InstrumentType) error
// }

// type someApplicationService struct {
// 	// ds   domain.SomeDomainService
// 	sites           SiteStorage
// 	instrumentTypes InstrumentTypeStorage
// 	events          chan domain.Event
// }

// // CommissionInstrument TODO
// func (service *someApplicationService) AddInstrument(
// 	siteID domain.SiteID,
// 	instrumentTypeID domain.InstrumentTypeID,
// ) error {
// 	site, err := service.sites.Load(siteID)
// 	if err != nil {
// 		return err
// 	}

// 	err = site.CommissionInstrument(instrumentID)

// }

// // AddInstrument TODO
// func (service *someApplicationService) AddInstrument(
// 	siteID domain.SiteID,
// 	typeID domain.InstrumentTypeID,
// 	iteration uint8,
// ) error {
// 	instrument, err := domain.NewInstrument(
// 		site,
// 		instrumentType,
// 		iteration,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	err = instrument.Add(service.events)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

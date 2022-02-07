package domain

// // Instrument TODO
// type Instrument struct {
// 	id               instrument.ID
// 	instrumentTypeID instrumentType.ID
// 	siteID           site.ID
// 	name             instrument.Name
// 	isActive         bool
// }

// // NewInstrument TODO
// func NewInstrument() instrument.Aggregate {
// 	return &Instrument{
// 		id:               NoInstrumentID{},
// 		instrumentTypeID: NoInstrumentTypeID{},
// 		siteID:           NoSiteID{},
// 		name:             instrument.Name(""),
// 	}
// }

// // Activate TODO
// func (ar *Instrument) Activate() error {
// 	if ar.isActive {
// 		return fmt.Errorf("instrument is already active")
// 	}

// 	ar.isActive = true

// 	return nil
// }

// // Deactivate TODO
// func (ar *Instrument) Deactivate() error {
// 	if !ar.isActive {
// 		return fmt.Errorf("instrument is already inactive")
// 	}

// 	ar.isActive = false

// 	return nil
// }

// // Relocate TODO
// func (ar *Instrument) Relocate(
// 	newSiteID site.ID,
// ) error {
// 	ar.siteID = newSiteID

// 	return nil
// }

// // Rename TODO
// func (ar *Instrument) Rename(
// 	newName instrument.Name,
// ) error {
// 	ar.name = newName

// 	return nil
// }

// // ChangeInstrumentType TODO
// func (ar *Instrument) ChangeInstrumentType(
// 	instrumentTypeID instrumentType.ID,
// ) error {
// 	ar.instrumentTypeID = instrumentTypeID

// 	return nil
// }

// // NoInstrumentID TODO
// type NoInstrumentID struct{}

// // String TODO
// func (NoInstrumentID) String() string {
// 	return "Currently not installed at any site."
// }

// // Equals TODO
// func (NoInstrumentID) Equals(
// 	otherInstrumentID instrument.ID,
// ) bool {
// 	var equals bool

// 	switch otherInstrumentID.(type) {
// 	case NoInstrumentID:
// 		equals = true
// 	default:
// 		equals = false
// 	}

// 	return equals
// }

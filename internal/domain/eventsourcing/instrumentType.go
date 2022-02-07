package eventsourcing

// // InstrumentType TODO
// type InstrumentType struct {
// 	id           domain.InstrumentTypeID
// 	name         domain.InstrumentTypeName
// 	abbreviation domain.InstrumentTypeAbbreviation
// }

// // Rename TODO
// func (instrumentType InstrumentType) Rename(
// 	newName domain.InstrumentTypeName,
// ) error {
// 	if len(newName) < 1 {
// 		return fmt.Errorf("name may not be an empty string")
// 	}

// 	instrumentType.name = newName

// 	return nil
// }

// // ChangeAbbreviation TODO
// func (instrumentType *InstrumentType) ChangeAbbreviation(
// 	newAbbreviation domain.InstrumentTypeAbbreviation,
// ) error {
// 	instrumentType.abbreviation = newAbbreviation

// 	return nil
// }

// // NoInstrumentTypeID TODO
// type NoInstrumentTypeID struct{}

// // String TODO
// func (NoInstrumentTypeID) String() string {
// 	return "Currently not installed at any site."
// }

// // Equals TODO
// func (NoInstrumentTypeID) Equals(
// 	otherInstrumentTypeID domain.InstrumentTypeID,
// ) bool {
// 	var equals bool

// 	switch otherInstrumentTypeID.(type) {
// 	case NoInstrumentTypeID:
// 		equals = true
// 	default:
// 		equals = false
// 	}

// 	return equals
// }

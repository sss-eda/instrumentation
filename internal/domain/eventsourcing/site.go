package eventsourcing

// // Site - This will be the aggregate root
// type Site struct {
// 	sync.Mutex
// 	id           domain.SiteID
// 	name         domain.SiteName
// 	abbreviation domain.SiteAbbreviation
// 	changes      []Event
// 	sequence     uint64
// }

// func (site *Site) mutate(
// 	event Event,
// ) {
// 	switch e := event.(type) {
// 	case SiteRenamedEvent:
// 		site.name = e.NewName
// 	case SiteAbbreviationChangedEvent:
// 		site.abbreviation = e.NewAbbreviation
// 	default:
// 		return
// 	}

// 	site.sequence++
// }

// func (site *Site) raise(
// 	event Event,
// ) {
// 	site.Lock()
// 	defer site.Unlock()

// 	site.mutate(event)
// 	site.changes = append(site.changes, event)
// }

// // NewSite TODO
// func NewSite(
// 	events []Event,
// ) (*Site, error) {
// 	site := &Site{
// 		Mutex:        sync.Mutex{},
// 		id:           NoSiteID{},
// 		name:         domain.SiteName(""),
// 		abbreviation: domain.SiteAbbreviation(""),
// 		changes:      []Event{},
// 		sequence:     0,
// 	}

// 	for _, event := range events {
// 		site.raise(event)
// 	}

// 	return site, nil
// }

// // Rename TODO
// func (site *Site) Rename(
// 	newName domain.SiteName,
// ) error {
// 	if newName == "" {
// 		return fmt.Errorf("name can't be an empty string")
// 	} else if newName == site.name {
// 		return fmt.Errorf("the new name is the same as the previous one")
// 	}

// 	site.raise(SiteRenamedEvent{
// 		NewName: newName,
// 	})

// 	return nil
// }

// // ChangeAbbreviation - TODO
// func (site *Site) ChangeAbbreviation(
// 	newAbbreviation domain.SiteAbbreviation,
// ) error {
// 	if newAbbreviation == "" {
// 		return fmt.Errorf("abbreviation can't be an empty string")
// 	} else if newAbbreviation == site.abbreviation {
// 		return fmt.Errorf("the new name is the same as the previous one")
// 	}

// 	site.raise(SiteAbbreviationChangedEvent{
// 		NewAbbreviation: newAbbreviation,
// 	})

// 	return nil
// }

// // NoSiteID TODO
// type NoSiteID struct{}

// // String TODO
// func (NoSiteID) String() string {
// 	return "Currently not installed at any site."
// }

// // Equals TODO
// func (NoSiteID) Equals(
// 	otherSiteID domain.SiteID,
// ) bool {
// 	var equals bool

// 	switch otherSiteID.(type) {
// 	case NoSiteID:
// 		equals = true
// 	default:
// 		equals = false
// 	}

// 	return equals
// }

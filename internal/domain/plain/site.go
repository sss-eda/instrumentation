package domain

import (
	"fmt"

	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// Site - This will be the aggregate root
type Site struct {
	id           site.ID
	name         site.Name
	abbreviation site.Abbreviation
	// changes      chan Event
}

// Rename TODO
func (ar Site) Rename(
	newName site.Name,
) error {
	if newName == "" {
		return fmt.Errorf("name can't be an empty string")
	} else if newName == ar.name {
		return fmt.Errorf("the new name is the same as the previous one")
	}

	ar.name = newName

	return nil
}

// ChangeAbbreviation - TODO
func (ar Site) ChangeAbbreviation(
	newAbbreviation site.Abbreviation,
) error {
	ar.abbreviation = newAbbreviation

	return nil
}

// NoSiteID TODO
type NoSiteID struct{}

// String TODO
func (NoSiteID) String() string {
	return "Currently not installed at any site."
}

// Equals TODO
func (NoSiteID) Equals(
	otherSiteID site.ID,
) bool {
	var equals bool

	switch otherSiteID.(type) {
	case NoSiteID:
		equals = true
	default:
		equals = false
	}

	return equals
}

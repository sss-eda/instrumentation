package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// ChangeSiteAbbreviationInput TODO
type ChangeSiteAbbreviationInput struct {
	SiteID              site.ID
	NewSiteAbbreviation site.Abbreviation
}

// ChangeSiteAbbreviationMutation TODO
func ChangeSiteAbbreviationMutation(
	storage site.Storage,
) (
	func(ChangeSiteAbbreviationInput) error,
	error,
) {
	return func(input ChangeSiteAbbreviationInput) error {
		aggregate, err := storage.Load(input.SiteID)
		if err != nil {
			return err
		}

		err = aggregate.ChangeAbbreviation(
			input.NewSiteAbbreviation,
		)
		if err != nil {
			return err
		}

		err = storage.Save(input.SiteID, aggregate)
		if err != nil {
			return err
		}

		return nil
	}, nil
}

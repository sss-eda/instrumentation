package domain

import (
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// RenameSiteInput TODO
type RenameSiteInput struct {
	SiteID      site.ID
	NewSiteName site.Name
}

// RenameSiteMutation TODO
func RenameSiteMutation(
	storage site.Storage,
) (
	func(RenameSiteInput) error,
	error,
) {
	return func(input RenameSiteInput) error {
		aggregate, err := storage.Load(input.SiteID)
		if err != nil {
			return err
		}

		err = aggregate.Rename(
			input.NewSiteName,
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

package domain

import (
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// RemoveSiteInput TODO
type RemoveSiteInput struct {
	SiteID site.ID
}

// RemoveSiteMutation TODO
func RemoveSiteMutation(
	storage site.Storage,
) (
	func(RemoveSiteInput) error,
	error,
) {
	return func(input RemoveSiteInput) error {
		aggregate, err := storage.Load(input.SiteID)
		if err != nil {
			return err
		}

		err = aggregate.Remove()
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

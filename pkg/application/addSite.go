package application

import (
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// AddSiteInput TODO
type AddSiteInput struct {
	SiteName         site.Name
	SiteAbbreviation site.Abbreviation
}

// AddSiteMutation TODO
func AddSiteMutation(
	factory site.Factory,
	storage site.Storage,
) (
	func(AddSiteInput) error,
	error,
) {
	return func(input AddSiteInput) error {
		id, aggregate, err := factory.New()
		if err != nil {
			return err
		}

		err = aggregate.Add(
			input.SiteName,
			input.SiteAbbreviation,
		)
		if err != nil {
			return err
		}

		err = storage.Save(id, aggregate)
		if err != nil {
			return err
		}

		return nil
	}, nil
}

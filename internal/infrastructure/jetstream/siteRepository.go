package jetstream

import "github.com/sss-eda/instrumentation/pkg/domain/site"

// SiteRepository TODO
type SiteRepository struct{}

// New TODO
func (repository *SiteRepository) New() (site.ID, site.Aggregate) {
	return nil, nil
}

// Load TODO
func (repository *SiteRepository) Load(
	id site.ID,
) (site.Aggregate, error) {
	return nil, nil
}

// Load TODO
func (repository *SiteRepository) Save(
	id site.ID,
	aggregate site.Aggregate,
) error {
	return nil
}

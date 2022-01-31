package jetstream

import (
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/instrumentation/internal/application/eventsourcing"
	"github.com/sss-eda/instrumentation/internal/domain"
)

// EventStore TODO
type EventStore struct {
	identifier identification.Service
	js         nats.JetStream
}

// NewID TODO
func (eventstore *EventStore) NewID() eventsourcing.AggregateID {
	id := eventstore.identifier.Next()
	return id
}

// Load TODO
func (repository *SiteRepository) Load(
	id domain.SiteID,
) (domain.Site, error) {
	site := repository.cache.Lookup(id)
	if site != nil {
		return site, nil
	}

	sub, err := repository.js.Subscribe(
		strings.Join([]string{
			"site",
			id.String(),
			">",
		}, "."),
		func(msg *nats.Msg) {},
	)
	if err != nil {
		return nil, err
	}

	return
}

// Save TODO
func (repository *SiteRepository) Save(
	id domain.SiteID,
	aggregate domain.Site,
) error {
	return nil
}

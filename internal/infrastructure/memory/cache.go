package memory

import (
	"github.com/sss-eda/instrumentation/internal/application/eventsourcing"
)

type cache struct {
	storage    eventsourcing.Storage
	aggregates map[eventsourcing.AggregateID]eventsourcing.Aggregate
}

// Cache TODO
func Cache(
	storage eventsourcing.Storage,
) (eventsourcing.Storage, error) {
	return &cache{
		storage: storage,
	}, nil
}

// Load TODO
func (c *cache) Load(
	id eventsourcing.AggregateID,
) (eventsourcing.Aggregate, error) {
	aggregate, ok := c.aggregates[id]
	if !ok {
		return nil, nil
	}

	return aggregate, nil
}

// Save TODO
func (c *cache) Save(
	id eventsourcing.AggregateID,
	aggregate eventsourcing.Aggregate,
) error {
	c.aggregates[id] = aggregate

	return nil
}

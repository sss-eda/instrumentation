package jetstream

import (
	"strings"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// Repository TODO
type Repository struct {
	js *nats.JetStream
}

// NewAggregate TODO
func (repository *Repository) NewAggregate() (AggregateID, Aggregate) {
	newID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	aggregate := factory.New(newID)

	return newID, aggregate
}

// Load TODO
func (repository *Repository) Load(
	aggregateID AggregateID,
	aggregateType AggregateType,
) (Aggregate, error) {
	aggregate := &Aggregate{}

	sub, err := repository.js.Subscribe(
		strings.Join([]string{
			"instrumentation",
			aggregateType.String(),
			aggregateID.String(),
			"events.>",
		}, "."),
		aggregate.Apply,
	)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	return aggregate, nil
}

// Save TODO
func (repository *Repository) Save(
	id AggregateID,
	aggregate Aggregate,
) error {
	return nil
}

package jetstream

import (
	"github.com/nats-io/nats.go"
	"github.com/sss-eda/instrumentation/pkg/application"
)

// RenameSite TODO
func RenameSite(
	mutation application.RenameSiteMutation,
) (nats.Handler, error) {
	return func(msg *nats.Msg) {

	}, nil
}

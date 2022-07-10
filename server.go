package instrumentation

import (
	"github.com/sss-eda/instrumentation/internal/servers/graphql"
	"github.com/sss-eda/instrumentation/internal/servers/openapi"
)

// Do we want the servers to be publicly available? No. If you need the server, run it... No?

// Server TODO
type Server interface {
	Run()
}

// NewServer TODO
func NewOpenAPIServer() (Server, error) {
	return openapi.NewServer()
}

// NewGraphQLServer TODO
func NewGraphQLServer() (Server, error) {
	return graphql.NewServer()
}

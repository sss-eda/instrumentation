package rest

import (
	"net/http"

	"github.com/sss-eda/instrumentation"
)

type API struct {
	Server instrumentation.Server
}

func (api *API) Serve() error {
	http.ListenAndServe(":8080", nil)
}

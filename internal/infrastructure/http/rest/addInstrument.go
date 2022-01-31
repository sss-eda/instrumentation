package rest

import (
	"encoding/json"
	"net/http"

	"github.com/sss-eda/instrumentation/internal/application"
	"github.com/sss-eda/instrumentation/pkg/domain/instrument"
	"github.com/sss-eda/instrumentation/pkg/domain/instrumentType"
	"github.com/sss-eda/instrumentation/pkg/domain/site"
)

// AddInstrument TODO
func AddInstrument(
	useCase application.InstrumentAdder,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := AddInstrumentRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = service.AddInstrument(
			request.Name,
			request.InstrumentTypeID,
			request.SiteID,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// AddInstrumentRequest TODO
type AddInstrumentRequest struct {
	Name             instrument.Name   `json:"name"`
	InstrumentTypeID instrumentType.ID `json:"instrumentTypeID"`
	SiteID           site.ID           `json:"siteID"`
}

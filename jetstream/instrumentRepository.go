package jetstream

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/instrumentation"
)

func GetInstrumentByID(js nats.JetStream) func(string, string) (*instrumentation.Instrument, error) {
	return func(id string, kind string) (*instrumentation.Instrument, error) {
		instrument := instrumentation.Instrument{}

		var event instrumentation.Event
		_, err := js.Subscribe(
			"instrumentation."+kind+"."+id+".events",
			func(msg *nats.Msg) {
				err := json.Unmarshal(msg.Data, &event)
				if err != nil {
					log.Println(err)
				}
				err = (&instrument).ApplyEvent(&event)
				if err != nil {
					log.Println(err)
				}
			},
		)
		if err != nil {
			return nil, err
		}

		return &instrument, nil
	}
}

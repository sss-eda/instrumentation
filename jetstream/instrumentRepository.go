package jetstream

import (
	natsio "github.com/nats-io/nats.go"
)

type Repository struct {
	js natsio.JetStream
}

func NewRepository(js natsio.JetStream) (*Repository, error) {
	return &Repository{
		js: js,
	}, nil
}

func (repository *Repository) CreateInstrument() error {
	repository.js.
}

// func GetInstrumentByID(js nats.JetStream) func(string, string) (*instrumentation.Instrument, error) {
// 	return func(id string, kind string) (*instrumentation.Instrument, error) {
// 		instrument := instrumentation.Instrument{}

// 		var event instrumentation.Event
// 		_, err := js.Subscribe(
// 			"instrumentation."+kind+"."+id+".events",
// 			func(msg *nats.Msg) {
// 				err := json.Unmarshal(msg.Data, &event)
// 				if err != nil {
// 					log.Println(err)
// 				}
// 				err = (&instrument).ApplyEvent(&event)
// 				if err != nil {
// 					log.Println(err)
// 				}
// 			},
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return &instrument, nil
// 	}
// }

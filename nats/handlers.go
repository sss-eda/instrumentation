package nats

import (
	"encoding/json"
	"log"
	"strings"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/instrumentation"
	"github.com/sss-eda/instrumentation/server"
)

// type HandlerFunc func(*natsio.Msg)

// func CommandSubscriptionhandlerFunc(service *server.Service, idFieldNumber int) func(*instrument.Command) {

// }

func CommandSubscriptionHandlerFunc(service *server.Service, idFieldNumber int) func(*natsio.Msg) {
	return func(msg *natsio.Msg) {
		command := instrumentation.Command("")
		id := strings.Split(msg.Subject, ".")[idFieldNumber]
		err := json.Unmarshal(msg.Data, &command)
		if err != nil {
			errMsg, err2 := json.Marshal(err)
			if err2 != nil {
				log.Printf("failed to parse error to json: %v", err2)
				err3 := msg.Respond([]byte(err.Error()))
				if err3 != nil {
					log.Printf("failed to respond with error: %v", err3)
					return
				}
				return
			}
			err2 = msg.Respond(errMsg)
			if err2 != nil {
				log.Printf("failed to respond with error: %v", err2)
				return
			}
		}
		err = service.SendCommandToInstrument(id, &command)
		if err != nil {
			errMsg, err2 := json.Marshal(err)
			if err2 != nil {
				log.Printf("failed to parse error to json: %v", err2)
				err3 := msg.Respond([]byte(err.Error()))
				if err3 != nil {
					log.Printf("failed to respond with error: %v", err3)
					return
				}
				return
			}
			err2 = msg.Respond(errMsg)
			if err2 != nil {
				log.Printf("failed to respond with error: %v", err2)
				return
			}
		}
	}
}

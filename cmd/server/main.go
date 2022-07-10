package main

import (
	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/instrumentation/api"
	"github.com/sss-eda/instrumentation/nats"
)

func main() {
	// Infrastructure
	nc, _ := natsio.Connect("wss://nats.sansa.dev/instrumentation")
	js, _ := nc.JetStream()

	instruments, _ := jetstream.NewInstrumentRepository(js)

	service, _ := api.New(instruments)

	js.Subscribe("instrumentation.*.*.commands", nats.CommandSubscriptionHandlerFunc(service, 2))
}

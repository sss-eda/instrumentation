package main

import (
	natsio "github.com/nats-io/nats.go"
)

func main() {
	// Infrastructure
	nc, _ := natsio.Connect("wss://js.sansa.dev/instrumentation")
	js, _ := nc.JetStream()

	js.Subscribe("instrumentation.*.*.commands", natsio.CommandsHandler())
}

package main

import "net/http"

func main() {
	http.HandleFunc(
		"/instrument/activate",
		rest.ActivateInstrument(
			instrumentation.ActivateInstrument(nats.InstrumentFactory{})),
	)

	http.HandleFunc(
		"" 
	)
}

// func main() {
// 	eventStore := memory.Cache(jetstream.EventStore{})

// 	app := application.New(eventStore)

// 	controller := rest.New(app)

// 	log.Fatal(controller.Run())
// }

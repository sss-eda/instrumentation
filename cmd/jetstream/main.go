package main

func main() {
	// nc, err := nats.Connect("nats://sansa.dev:4222")
	// js, err := nc.jetstream()

	// sub1, err := js.Subscribe(
	// 	"INSTRUMENTATION.COMMANDS.instrument.*.activate",
	// 	jetstream.ActivateInstrumentAdapter(
	// 		instrumentation.ActivateInstrument(
	// 			jetstream.EventStore{js},

	// 		),
	// 	),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer sub1.Unsubscribe()

	// http.HandleFunc(
	// 	"/instrument/activate",
	// 	rest.ActivateInstrument(
	// 		instrumentation.ActivateInstrument(nats.InstrumentFactory{})),
	// )

	// http.HandleFunc(
	// 	""
	// )
}

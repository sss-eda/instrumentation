package main

import (
	"context"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

// It's possible that all of this belongs in the "configuration" domain...?
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	siteID := os.Getenv("SITE_ID")

	nc, err := nats.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	site := site.New(siteID, js) // This site is not an aggregate. It is
	// simply a projection of the event stream for site configuration events.

	for _, instrument := range site.Instruments() {
		go instrument.Run(ctx)
	}
	defer cancel()

	<-ctx.Done()
	log.Fatal(ctx.Err())
}

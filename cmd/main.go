package main

import (
	"log"

	"github.com/sss-eda/instrumentation/internal/application"
	"github.com/sss-eda/instrumentation/internal/infrastructure/memory"
)

func main() {
	eventStore := memory.Cache(jetstream.EventStore{})

	app := application.New(instrumentRepo, siteRepo, instrumentTypeRepo)
	log.Fatal(app.Run())
}

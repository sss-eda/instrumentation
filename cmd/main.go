package main

import (
	"log"

	"github.com/sss-eda/instrumentation/internal/application"
)

func main() {
	instrumentRepo := jetstream.NewInstrumentRepository()
	siteRepo := jetstream.NewSiteRepository()
	instrumentTypeRepo := jetstream.NewInstrumentTypeRepository()

	app := application.New(instrumentRepo, siteRepo, instrumentTypeRepo)
	log.Fatal(app.Run())
}

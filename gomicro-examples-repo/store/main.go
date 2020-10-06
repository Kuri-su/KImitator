package main

import (
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.store-test"),
		micro.Version("latest"),
	)

	service.Init()

	// write to the store
	record := &store.Record{Key: "foo", Value: []byte("bar")}
	if err := service.Options().Store.Write(record); err != nil {
		logger.Fatalf("Error writing to the store: %v", err)
	}
	logger.Infof("Wrote value to the store")

	// read from the store
	go func() {
		ticker := time.NewTicker(time.Second * 15)

		for {
			<-ticker.C

			recs, err := service.Options().Store.Read("foo")
			if err != nil {
				logger.Errorf("Error reading from store: %v", err)
				continue
			}
			if len(recs) == 0 {
				logger.Errorf("No results returned from store")
				continue
			}
			logger.Infof("Got a value from the store: %v", string(recs[0].Value))
		}
	}()

	// run the service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

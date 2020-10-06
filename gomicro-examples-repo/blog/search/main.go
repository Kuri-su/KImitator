package main

import (
	"github.com/micro/examples/blog/search/handler"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	search "github.com/micro/examples/blog/search/proto/search"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.search"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	search.RegisterSearchServiceHandler(service.Server(), new(handler.Search))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

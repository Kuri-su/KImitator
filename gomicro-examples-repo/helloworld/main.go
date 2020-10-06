package main

import (
	"github.com/micro/examples/helloworld/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"

	pb "github.com/micro/examples/helloworld/proto"
)

func main() {
	// New Service
	helloworld := micro.NewService(
		micro.Name("helloworld"),
		micro.Version("latest"),
	)

	// Initialise service
	helloworld.Init()

	// Register Handler
	pb.RegisterHelloworldHandler(helloworld.Server(), new(handler.Helloworld))

	// Run service
	if err := helloworld.Run(); err != nil {
		logger.Fatal(err)
	}
}

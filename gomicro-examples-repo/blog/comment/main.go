package main

import (
	"github.com/micro/examples/blog/comment/handler"
	"github.com/micro/examples/blog/comment/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	comment "github.com/micro/examples/blog/comment/proto/comment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.comment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	comment.RegisterCommentHandler(service.Server(), new(handler.Comment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.comment", service.Server(), new(subscriber.Comment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

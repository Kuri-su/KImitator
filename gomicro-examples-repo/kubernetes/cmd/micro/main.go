package main

import (
	"os"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	cli "github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/server"
	srv "github.com/micro/go-micro/server/grpc"
	bkr "github.com/micro/go-plugins/broker/grpc"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	_ "github.com/micro/go-plugins/client/selector/static"

	// disable namespace by default
	_	"github.com/micro/go-micro/api"
)

func main() {

	// set values for registry/selector
	os.Setenv("MICRO_REGISTRY", "kubernetes")
	os.Setenv("MICRO_SELECTOR", "static")

	// setup broker/client/server
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()

	// init command
	cmd.Init()
}

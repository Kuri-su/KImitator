package _enter

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"log"
	"time"
)

func main() {
	service := micro.NewService(
		micro.Name("name"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)

	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"localhost:8500",
		}
	})
	// Init will parse the command line flags.
	service.Init(
		micro.Registry(reg),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Fatalln("start service failed")
	}
}
